import { api } from "@/lib/api";
import { Message, MetricsResponse } from "@/types/common";
import { useCallback, useEffect, useState } from "react";

export function useMetrics(refreshInterval: number = 1000){
    const [metrics, setMetrics] = useState<MetricsResponse | null>(null)
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<Error | null>(null);

    useEffect(() => {
        const fetchMetrics = async () => {
            try{
                const data = await api.getMetrics();
                setMetrics(data);
                setError(null);
            }
            catch(err){
                setError(err as Error);
                console.error("Failed to fetch metrics:", err)
            } finally{
                setLoading(false);
            }
        }
        fetchMetrics();
        const interval = setInterval(fetchMetrics, refreshInterval + 500);

        return () => clearInterval(interval);
    }, [refreshInterval]);

    const refresh = useCallback(async () => {
        setLoading(true);
        try{
            const data = await api.getMetrics();
            setMetrics(data);
            setError(null);
        } catch (err) {
            setError(err as Error);
        } finally {
            setLoading(false);
        }
    }, []);

    const resetMetrics = useCallback(async () => {
        try{
            await api.resetMetrics();
            await refresh();
        } catch (err) {
            setError (err as Error);
            console.error("Failed to reset metrics:", err)
        }
    }, [refresh])

    return { metrics, loading, error, refresh, resetMetrics };
}


