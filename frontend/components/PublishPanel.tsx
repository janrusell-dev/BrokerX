"use client";

import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { usePublish } from "@/hooks/usePublish";

export function PublishPanel() {
  const [topic, setTopic] = useState("orders");
  const [sender, setSender] = useState("frontend-dashboard");
  const [message, setMessage] = useState("");
  const { publish, loading } = usePublish();
  const [status, setStatus] = useState("");

  const handleSend = async () => {
    if (!message.trim()) return;

    try {
      const response = await publish({
        topic,
        sender,
        payload: {
          message: message,
          timestamp: new Date().toISOString(),
        },
      });

      setStatus(`✓ Sent (${response.latency}ms)`);
      setMessage("");
      setTimeout(() => setStatus(""), 2000);
    } catch {
      setStatus("✗ Failed");
      setTimeout(() => setStatus(""), 2000);
    }
  };

  return (
    <Card className="bg-neutral-900 border-neutral-800 text-white">
      <CardHeader>
        <CardTitle>Publish Message</CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        <Input
          placeholder="Topic"
          value={topic}
          onChange={(e) => setTopic(e.target.value)}
          className="bg-neutral-800 border-neutral-700"
        />
        <Input
          placeholder="Sender"
          value={sender}
          onChange={(e) => setSender(e.target.value)}
          className="bg-neutral-800 border-neutral-700"
        />
        <Input
          placeholder="Message..."
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && handleSend()}
          className="bg-neutral-800 border-neutral-700"
        />
        <Button
          className="w-full bg-emerald-500 hover:bg-emerald-600"
          onClick={handleSend}
          disabled={loading || !message.trim()}
        >
          {loading ? "Sending..." : "Send"}
        </Button>
        {status && (
          <p className="text-sm text-center text-emerald-400">{status}</p>
        )}
      </CardContent>
    </Card>
  );
}
