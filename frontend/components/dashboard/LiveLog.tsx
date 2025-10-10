'use client'

import { useEffect, useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Badge } from "../ui/badge";

export function LiveLog(){
    const [messages, setMessages] = useState<any[]>([]);

    useEffect(() => {
        const ws = new WebSocket("ws://localhost:8080/subscribe?topic=orders");
        ws.onmessage = (e) => {
            const msg = JSON.parse(e.data);
            setMessages((prev) => [msg, ...prev.slice(0, 10)]);
        };
        return () => ws.close();
    }, []);
     return (
    <Card className="bg-neutral-900 border-neutral-800 text-white">
      <CardHeader>
        <CardTitle>Live Messages</CardTitle>
      </CardHeader>
      <CardContent className="space-y-2 max-h-80 overflow-y-auto">
        {messages.map((m, i) => (
          <div key={i} className="border-b border-neutral-800 pb-2">
            <div className="flex items-center justify-between">
              <span className="text-emerald-400 font-medium">{m.sender}</span>
              <Badge variant="secondary">{m.topic}</Badge>
            </div>
            <p className="text-gray-300">{m.payload}</p>
            <p className="text-xs text-gray-500">{new Date(m.timestamp).toLocaleTimeString()}</p>
          </div>
        ))}
      </CardContent>
    </Card>
  );
}