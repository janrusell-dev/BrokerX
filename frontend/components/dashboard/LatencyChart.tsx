"use client";

import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  CartesianGrid,
} from "recharts";

type LatencyData = {
  time: string;
  latency: number;
};

export function LatencyChart({ data }: { data: LatencyData[] }) {
  return (
    <div className="bg-neutral-900 p-4 rounded-2xl shadow-lg">
      <h2 className="text-lg font-semibold text-emerald-400 mb-2">
        Latency Over Time (ms)
      </h2>
      <ResponsiveContainer width="100%" height={250}>
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" stroke="#333" />
          <XAxis dataKey="time" stroke="#aaa" />
          <YAxis stroke="#aaa" />
          <Tooltip />
          <Line
            type="monotone"
            dataKey="latency"
            stroke="#10b981"
            strokeWidth={2}
            dot={false}
            isAnimationActive={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
