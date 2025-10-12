
import { api } from "@/lib/api";
import { useEffect, useState } from "react";

export function useHealth(checkInterval: number = 5000) {
  const [isHealthy, setIsHealthy] = useState(false);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkHealth = async () => {
      try {
        await api.health()
        setIsHealthy(true);
      } catch {
        setIsHealthy(false);
      } finally {
        setLoading(false);
      }
    };

    checkHealth();
    const interval = setInterval(checkHealth, checkInterval);

    return () => clearInterval(interval);
  }, [checkInterval]);

  return { isHealthy, loading };
}