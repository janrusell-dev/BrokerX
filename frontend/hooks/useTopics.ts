import { api } from "@/lib/api";
import { useCallback, useEffect, useState } from "react";

export function useTopics(refreshInterval?: number){
    const [topics, setTopics] = useState<string[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<Error | null>(null);

    const fetchTopics = useCallback(async () => {
        try{
            const data = await api.getTopics();
            setTopics(data.topics);
            setError(null);
        } catch(err) {
            setError(err as Error);
            console.error("Failed to fetch topics:", err);
        } finally{
            setLoading(false);
        }
    }, []);
    
     useEffect(() => {
    fetchTopics();

    if (refreshInterval) {
      const interval = setInterval(fetchTopics, refreshInterval);
      return () => clearInterval(interval);
    }
  }, [refreshInterval, fetchTopics]);

  return { topics, loading, error, refresh: fetchTopics };
}