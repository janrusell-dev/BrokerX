'use client'

import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { Input } from "./ui/input";
import { Button } from "./ui/button";

export function PublishPanel() {
    const [topic, setTopic] = useState("orders");
    const [sender, setSender] = useState("frontend-dashboard");
    const [payload, setPayload] = useState("");

    const sendMessage = async () => {
         await fetch("http://localhost:8080/publish", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ topic, sender, payload }),
    });
      setPayload("");
    }

     return (
    <Card className="bg-neutral-900 border-neutral-800 text-white">
      <CardHeader>
        <CardTitle>Publish Message</CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        <Input placeholder="Topic" value={topic} onChange={(e) => setTopic(e.target.value)} />
        <Input placeholder="Sender" value={sender} onChange={(e) => setSender(e.target.value)} />
        <Input placeholder="Message..." value={payload} onChange={(e) => setPayload(e.target.value)} />
        <Button className="w-full bg-emerald-500 hover:bg-emerald-600" onClick={sendMessage}>
          Send
        </Button>
      </CardContent>
    </Card>
  );
} 