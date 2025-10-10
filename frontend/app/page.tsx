"use client";

import { LiveLog } from "@/components/dashboard/LiveLog";
import { MetricCard } from "@/components/dashboard/MetricCard";
import { PublishPanel } from "@/components/PublishPanel";
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
  const [metrics, setMetrics] = useState<any>(null);
  const [chartData, setChartData] = useState<any[]>([]);
  const [latencyData, setLatencyData] = useState<any[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetch("http://localhost:8080/metrics");
      const data = await res.json();
      setMetrics(data);

      // Bar chart for topic message counts
      const topics = Object.entries(data.messages_per_topic || {}).map(
        ([topic, count]) => ({
          topic,
          count,
        })
      );
      setChartData(topics);

      // Line chart for latency over time
      setLatencyData((prev) => {
        const newData = [
          ...prev.slice(-19),
          { time: new Date().toLocaleTimeString(), latency: data.avg_latency },
        ];
        return newData;
      });
    };

    fetchData();
    const interval = setInterval(fetchData, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <main className="min-h-screen bg-gradient-to-b from-neutral-950 via-neutral-900 to-black text-white p-6 space-y-8">
      <h1 className="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-cyan-400">
        BrokerX Realtime Dashboard
      </h1>

      {metrics && (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <MetricCard title="Total Messages" value={metrics.total_messages} />
          <MetricCard
            title="Avg Latency (ms)"
            value={metrics.avg_latency.toFixed(2)}
          />
          <MetricCard
            title="Active Subscribers"
            value={metrics.active_subscribers}
          />
          <MetricCard
            title="Topics"
            value={Object.keys(metrics?.messages_per_topic || {}).length}
          />
        </div>
      )}

      <div className="grid md:grid-cols-3 gap-6">
        <div className="md:col-span-2 space-y-6">
          {/* Topic Message Chart */}
          <div className="bg-neutral-900 p-4 rounded-2xl shadow-lg">
            <h2 className="text-lg font-semibold mb-2 text-emerald-300">
              Messages per Topic
            </h2>
            <ResponsiveContainer width="100%" height={250}>
              <BarChart data={chartData}>
                <CartesianGrid strokeDasharray="3 3" stroke="#333" />
                <XAxis dataKey="topic" stroke="#aaa" />
                <YAxis stroke="#aaa" />
                <Tooltip
                  contentStyle={{
                    backgroundColor: "#111",
                    border: "1px solid #333",
                  }}
                />
                <Bar dataKey="count" fill="#10b981" radius={[6, 6, 0, 0]} />
              </BarChart>
            </ResponsiveContainer>
          </div>

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
                  }}
                />
                <Line
                  type="monotone"
                  dataKey="latency"
                  stroke="#06b6d4"
                  strokeWidth={2}
                  dot={false}
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
