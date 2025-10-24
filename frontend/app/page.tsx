"use client";

import { LiveLog } from "@/components/dashboard/LiveLog";
import { MetricCard } from "@/components/dashboard/MetricCard";
import { TopicChart } from "@/components/dashboard/TopicChart";
import { PublishPanel } from "@/components/PublishPanel";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { useMetrics } from "@/hooks/useMetrics";
import { useEffect, useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  CartesianGrid,
  BarChart,
  Bar,
} from "recharts";

export default function Home() {
  const { metrics, loading, error, resetMetrics } = useMetrics(1000);
  const [chartData, setChartData] = useState<any[]>([]);
  const [latencyData, setLatencyData] = useState<any[]>([]);

  // Update charts when metrics change
  useEffect(() => {
    if (!metrics) return;

    // Bar chart for topic message counts
    const topicMetrics = metrics.topicMetrics || {};
    const topics = Object.entries(topicMetrics).map(
      ([topic, count]: [string, any]) => ({
        topic,
        count,
      })
    );
    setChartData(topics);

    // Line chart for latency over time
    if (metrics.latencyHistory && metrics.latencyHistory.length > 0) {
      const formattedLatency = metrics.latencyHistory
        .slice(-20)
        .map((point: any) => ({
          time: new Date(point.timestamp).toLocaleTimeString(),
          latency: point.latency,
        }));
      setLatencyData(formattedLatency);
    }
  }, [metrics]);

  if (loading && !metrics) {
    return (
      <main className="min-h-screen bg-gradient-to-b from-neutral-950 via-neutral-900 to-black text-white p-6 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-400 mx-auto mb-4"></div>
          <p className="text-gray-400">Loading BrokerX Dashboard...</p>
        </div>
      </main>
    );
  }

  if (error) {
    return (
      <main className="min-h-screen bg-gradient-to-b from-neutral-950 via-neutral-900 to-black text-white p-6 flex items-center justify-center">
        <div className="text-center">
          <p className="text-red-400 mb-2">
            Failed to connect to BrokerX backend
          </p>
          <p className="text-gray-500 text-sm">{error.message}</p>
        </div>
      </main>
    );
  }

  return (
    <main className="min-h-screen bg-gradient-to-b from-neutral-950 via-neutral-900 to-black text-white p-6 space-y-8">
      <div className="flex items-center justify-between">
        <h1 className="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-cyan-400">
          BrokerX Dashboard
        </h1>
        <div className="flex items-center gap-3">
          <Button
            onClick={resetMetrics}
            className="px-3 py-2 text-sm bg-red-600 hover:bg-red-700 rounded-lg font-medium transition"
          >
            Reset Metrics
          </Button>
          {metrics && (
            <Badge
              variant="outline"
              className="text-emerald-400 border-emerald-400"
            >
              Uptime: {Math.floor(metrics.uptime || 0)}s
            </Badge>
          )}
        </div>
      </div>

      {metrics && (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <MetricCard
            title="Total Messages"
            value={metrics.totalMessages || 0}
          />
          <MetricCard
            title="Avg Latency (ms)"
            value={(metrics.avgLatency || 0).toFixed(2)}
          />
          <MetricCard
            title="Active Subscribers"
            value={metrics.activeSubscribers || 0}
          />
          <MetricCard
            title="Message Rate"
            value={`${(metrics?.messageRate || 0).toFixed(1)}/s`}
          />
        </div>
      )}
      <div className="grid md:grid-cols-3 gap-6">
        <div className="md:col-span-2 space-y-6">
          {/* Topic Message Chart */}
          <TopicChart data={chartData} />

          {/* Latency Trend Chart */}
          <div className="bg-neutral-900 p-4 rounded-2xl shadow-lg">
            <h2 className="text-lg font-semibold mb-2 text-cyan-300">
              Latency Over Time
            </h2>
            <ResponsiveContainer width="100%" height={250}>
              <LineChart data={latencyData}>
                <CartesianGrid strokeDasharray="3 3" stroke="#333" />
                <XAxis dataKey="time" stroke="#aaa" />
                <YAxis stroke="#aaa" />
                <Tooltip
                  contentStyle={{
                    backgroundColor: "#111",
                    border: "1px solid #333",
                    borderRadius: "8px",
                  }}
                />
                <Line
                  type="monotone"
                  dataKey="latency"
                  stroke="#06b6d4"
                  strokeWidth={2}
                  dot={false}
                  isAnimationActive={false}
                />
              </LineChart>
            </ResponsiveContainer>
          </div>

          {/* Live Logs */}
          <LiveLog />
        </div>

        {/* Publish Panel */}
        <PublishPanel />
      </div>
    </main>
  );
}
