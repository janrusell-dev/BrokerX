import { LatencyPoint, Message, MetricsResponse, PublishRequest, PublishResponse, TopicInfo } from "@/types/common";
import { url } from "inspector";

// Centralized API service for BrokerX backend
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";
const WS_BASE_URL = process.env.NEXT_PUBLIC_WS_URL || "ws://localhost:8080";


export type WebSocketMessageCallback = (message: Message) => void;
export type WebSocketConnectCallback = () => void;
export type WebSocketDisconnectCallback = () => void;
export type WebSocketErrorCallback = (error: string) => void;

class BrokerXAPI {
    private baseUrl: string;
    private wsUrl: string;
    
    constructor(baseUrl: string = API_BASE_URL, wsUrl: string = WS_BASE_URL){
        this.baseUrl = baseUrl;
        this.wsUrl = wsUrl;
    }

    async publish(request: PublishRequest) : Promise<PublishResponse> {
        const response = await fetch(`${this.baseUrl}/publish`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(request)
        })

         if (!response.ok) {
      throw new Error(`Failed to publish message: ${response.statusText}`);
    }

    return response.json();
    }

    async getMetrics(): Promise<MetricsResponse> {
        const response = await fetch(`${this.baseUrl}/metrics`);

        if (!response.ok) {
            throw new Error(`Failed to fetch metrics: ${response.statusText}`)
        }

        return response.json();
    }

    async getLatencyHistory(): Promise<{ history: LatencyPoint[] }> {
        const response = await fetch(`${this.baseUrl}/metrics/latency`);

        if (!response.ok){
            throw new Error(`Failed to fetch latency history: ${response.statusText}`)
        }

        return response.json();
    }

    async resetMetrics(): Promise<{ status: string; message: string;}> {
        const response = await fetch(`${this.baseUrl}/metrics/reset`, {
            method: "POST",
        });

        if (!response.ok){
            throw new Error(`Failed to reset metrics: ${response.statusText}`)
        }

        return response.json();
    }

    async getTopics(): Promise<{ topics: string[] }> {
        const response = await fetch(`${this.baseUrl}/topics`);

        if (!response.ok){
            throw new Error(`Failed to fetch topics: ${response.statusText}`)
        }
        return response.json();
    }
    
    async getTopicInfo(topic: string): Promise<TopicInfo>{
        const response = await fetch(`${this.baseUrl}/topics/${topic}`);

        if (!response.ok){
            throw new Error(`Failed to fetch topic info ${response.statusText}`);
        }

        return response.json()
    }

    async getAllTopicsInfo(): Promise<{ topics: TopicInfo[] }>{
        const response = await fetch(`${this.baseUrl}/topics/info/all`);

        if (!response.ok){
            throw new Error(`Failed to fetch all topics info: ${response.statusText}`)
        }

        return response.json()
    }

    async health(): Promise<{ status: string; service: string}> {
        const response = await fetch(`${this.baseUrl}/health`);

        if (!response.ok){
            throw new Error(`Health check failed: ${response.statusText}`);
        }

        return response.json();
    }

   subscribe(
    topic: string,
    onMessage: (message: Message) => void,
    onConnect?: () => void,
    onDisconnect?: () => void,
    onError?: (error: string) => void
  ): WebSocket {
    const wsUrl = `${this.wsUrl}/subscribe?topic=${encodeURIComponent(topic)}`;
    
    let ws: WebSocket;
    
    try {
      ws = new WebSocket(wsUrl);
    } catch (error) {
      const errorMsg = `Failed to create WebSocket connection: ${error}`;
      console.error(`[WebSocket] ${errorMsg}`);
      onError?.(errorMsg);
      throw error;
    }

    let isOpen = false;

    ws.onopen = () => {
      isOpen = true;
      console.log(`[WebSocket] ✓ Connected to topic: ${topic}`);
      onConnect?.();
    };

    ws.onmessage = (event) => {
      if (!isOpen) return; // Ignore messages if connection is closing
      
      try {
        const message = JSON.parse(event.data);
        
        // Skip connection confirmation messages
        if (message.type === "connected") {
          console.log(`[WebSocket] Connection confirmed for topic: ${topic}`);
          return;
        }

        onMessage(message);
      } catch (error) {
        console.error("[WebSocket] Failed to parse message:", error);
        onError?.("Failed to parse message");
      }
    };

    ws.onclose = (event) => {
      isOpen = false;
      if (event.wasClean) {
        console.log(`[WebSocket] ✓ Cleanly disconnected from topic: ${topic} (code: ${event.code})`);
      } else {
        console.warn(`[WebSocket] ⚠ Connection lost for topic: ${topic} (code: ${event.code})`);
      }
      onDisconnect?.();
    };

    ws.onerror = (event) => {
      isOpen = false;
      // WebSocket errors don't provide much detail in the browser
      // Only log if we never successfully connected
      const errorMsg = `WebSocket connection error for topic: ${topic}. Check if backend is running at ${this.wsUrl}`;
      
      // Don't spam console with errors during hot reload
      if (ws.readyState === WebSocket.CONNECTING) {
        console.warn(`[WebSocket] Connection failed for topic: ${topic}`);
      }
      
      onError?.(errorMsg);
    };

    return ws;
  }

    createSubscription(topic: string){
        return{
            topic,
            url: `${this.wsUrl}/subscribe?topic=${topic}`,
        }
    }
}

export const api = new BrokerXAPI();

export default BrokerXAPI;