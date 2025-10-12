import { api } from "@/lib/api";
import { PublishRequest } from "@/types/common";
import { useCallback, useState } from "react";

export function usePublish() {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<Error | null>(null);
    const [lastResponse, setLastResponse] = useState<any>(null);

    const publish = useCallback(async (request: PublishRequest) => {
        setLoading(true);
        setError(null);

        try{
            const response = await api.publish(request);
            setLastResponse(response);
            return response;
        }
        catch (err){
            const error = err as Error;
            setError(error);
            throw error;
        } finally {
            setLoading(false);
        }
    }, []);
    return { publish, loading, error, lastResponse };
}