import { api } from "@/lib/api";
import { useCallback, useState } from "react";

export function useResetMetrics() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const reset = useCallback(async () => {
    setLoading(true);
    setError(null);

    try {
      await api.resetMetrics()
    } catch (err) {
      setError(err as Error);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  return { reset, loading, error };
}