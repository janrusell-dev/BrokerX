import { api } from "@/lib/api";
import { MetricsResponse } from "@/types/common";
import { useCallback, useEffect, useState } from "react";

export function useMetrics(refreshInterval: number = 1000) {
  const [metrics, setMetrics] = useState<MetricsResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  const fetchMetrics = useCallback(async () => {
    try {
      const data = await api.getMetrics();
      setMetrics(data);
      setError(null);
    } catch (err) {
      setError(err as Error);
      console.error("Failed to fetch metrics:", err);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchMetrics(); // initial fetch
    const interval = setInterval(fetchMetrics, refreshInterval);
    return () => clearInterval(interval);
  }, [fetchMetrics, refreshInterval]);

  const refresh = useCallback(async () => {
    await fetchMetrics();
  }, [fetchMetrics]);

  const resetMetrics = useCallback(async () => {
    try {
      await api.resetMetrics();
      await refresh();
    } catch (err) {
      setError(err as Error);
      console.error("Failed to reset metrics:", err);
    }
  }, [refresh]);

  return { metrics, loading, error, refresh, resetMetrics };
}
