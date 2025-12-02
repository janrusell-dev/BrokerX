import { api } from "@/lib/api";
import { TopicInfo } from "@/types/common";
import { useEffect, useState } from "react";

export function useTopicInfo(topic: string | null) {
  const [info, setInfo] = useState<TopicInfo | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    if (!topic) {
      setInfo(null);
      return;
    }

    const fetchInfo = async () => {
      setLoading(true);
      try {
        const data = await api.getTopicInfo(topic)
        setInfo(data);
        setError(null);
      } catch (err) {
        setError(err as Error);
      } finally {
        setLoading(false);
      }
    };

    fetchInfo();
  }, [topic]);

  return { info, loading, error };
}