import { api } from "@/lib/api";
import { SimulatorStatus } from "@/types/common";
import { useEffect, useState } from "react";

export function useSimulator() {
  const [isSimRunning, setIsSimRunning] = useState<boolean | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchStatus = async () => {
      try {
        const res: SimulatorStatus = await api.getSimulatorStatus();
        setIsSimRunning(res.running);
      } catch (err) {
        if (err instanceof Error) {
          setError(err);
        } else {
          setError(new Error("Failed to fetch simulator status"));
        }
      }
    };
    fetchStatus();
  }, []);

  const startSimulator = async () => {
    setLoading(true);
    setError(null);
    try {
      await api.startSimulator();
      setIsSimRunning(true);
    } catch (err) {
      if (err instanceof Error) {
        setError(err);
      } else {
        setError(new Error("Failed to start simulator"));
      }
    } finally {
      setLoading(false);
    }
  };

  const stopSimulator = async () => {
    setLoading(true);
    setError(null);
    try {
      await api.stopSimulator();
      setIsSimRunning(false);
    } catch (err) {
      if (err instanceof Error) {
        setError(err);
      } else {
        setError(new Error("Failed to stop simulator"));
      }
    } finally {
      setLoading(false);
    }
  };

  const toggleSimulator = async () => {
    if (isSimRunning === null) return;

    if (isSimRunning) {
      await stopSimulator();
    } else {
      await startSimulator();
    }
  };

  return { isSimRunning, loading, error, startSimulator, stopSimulator, toggleSimulator };
}
