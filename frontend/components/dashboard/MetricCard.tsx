import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";

interface MetricCardProps {
  title: string;
  value: number | string;
}

export function MetricCard({ title, value }: MetricCardProps) {
  return (
    <Card className="bg-neutral-900 border-neutral-800 text-white">
      <CardHeader>
        <CardTitle className="text-sm text-gray-400">{title}</CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-2xl font-bold text-emerald-400">{value}</p>
      </CardContent>
    </Card>
  );
}
