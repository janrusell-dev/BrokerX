import { api } from "@/lib/api";
import { Message } from "@/types/common";
import { useCallback, useEffect, useRef, useState } from "react";

export function useSubscription(topic: string) {
  const [messages, setMessages] = useState<Message[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const wsRef = useRef<WebSocket | null>(null);
  const reconnectTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const mountedRef = useRef(true);

  useEffect(() => {
    if (!topic) return;

    mountedRef.current = true;
    let reconnectAttempts = 0;
    const maxReconnectAttempts = 5;

    const connect = () => {
      // Don't try to connect if we're unmounting or already connected
      if (!mountedRef.current || (wsRef.current && wsRef.current.readyState === WebSocket.OPEN)) {
        return;
      }

      // Clear any existing connection
      if (wsRef.current) {
        wsRef.current.close();
        wsRef.current = null;
      }

      try {
        const ws = api.subscribe(
          topic,
          (message) => {
            if (mountedRef.current) {
              setMessages((prev) => [message, ...prev.slice(0, 49)]);
            }
          },
          () => {
            if (mountedRef.current) {
              setIsConnected(true);
              setError(null);
              reconnectAttempts = 0;
              console.log(`✅ Connected to topic: ${topic}`);
            }
          },
          () => {
            if (mountedRef.current) {
              setIsConnected(false);
              console.log(`⚠️ Disconnected from topic: ${topic}`);
              
              // Auto-reconnect after a delay (unless we're unmounting)
              if (reconnectAttempts < maxReconnectAttempts) {
                reconnectAttempts++;
                const delay = Math.min(1000 * reconnectAttempts, 5000);
                console.log(`Reconnecting in ${delay}ms... (attempt ${reconnectAttempts}/${maxReconnectAttempts})`);
                
                reconnectTimeoutRef.current = setTimeout(() => {
                  if (mountedRef.current) {
                    connect();
                  }
                }, delay);
              } else {
                setError("Connection lost. Refresh the page to reconnect.");
              }
            }
          },
          (errorMsg) => {
            if (mountedRef.current) {
              setError(errorMsg);
              setIsConnected(false);
              console.error(`❌ WebSocket error: ${errorMsg}`);
            }
          }
        );

        wsRef.current = ws;
      } catch (err) {
        if (mountedRef.current) {
          setError(err instanceof Error ? err.message : "Failed to connect");
          setIsConnected(false);
        }
      }
    };

    // Initial connection with small delay to avoid hot reload race conditions
    const initialTimeout = setTimeout(connect, 100);

    return () => {
      mountedRef.current = false;
      
      // Clear any reconnection timeouts
      if (reconnectTimeoutRef.current) {
        clearTimeout(reconnectTimeoutRef.current);
      }
      
      // Clear initial connection timeout
      clearTimeout(initialTimeout);
      
      // Close WebSocket connection
      if (wsRef.current) {
        // Remove event listeners to prevent callbacks during cleanup
        wsRef.current.onopen = null;
        wsRef.current.onmessage = null;
        wsRef.current.onerror = null;
        wsRef.current.onclose = null;
        
        if (wsRef.current.readyState === WebSocket.OPEN || wsRef.current.readyState === WebSocket.CONNECTING) {
          wsRef.current.close();
        }
        wsRef.current = null;
      }
    };
  }, [topic]);

  const clearMessages = useCallback(() => {
    setMessages([]);
  }, []);

  return { messages, isConnected, error, clearMessages };
}