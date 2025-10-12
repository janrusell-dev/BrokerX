'use client'

import { useEffect, useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Badge } from "../ui/badge";
import { useSubscription } from "@/hooks/useSubscription";

export function LiveLog(){
    const { messages, isConnected, error} = useSubscription("orders")

    return (
    <Card className="bg-neutral-900 border-neutral-800 text-white">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle>Live Messages (Orders)</CardTitle>
        <div className="flex items-center gap-2">
          <div className={`w-2 h-2 rounded-full ${isConnected ? 'bg-emerald-400' : 'bg-red-400'}`} />
          <span className="text-xs text-gray-400">
            {isConnected ? 'Connected' : 'Disconnected'}
          </span>
        </div>
      </CardHeader>
      <CardContent className="space-y-2 max-h-80 overflow-y-auto">
        { error && (
          <div className="bg-red-900/20 border border-red-500 text-red-400 p-3 rounded text-sm">
            <p className="font-semibold mb-1">Connection Error</p>
            <p className="text-xs">{error}</p>
          </div>
        )}
        {!error && messages.length === 0 ? (
          <p className="text-gray-500 text-center py-4">Waiting for messages...</p>
        ) : (
        messages.map((m, i) => (
          <div key={i} className="border-b border-neutral-800 pb-2">
            <div className="flex items-center justify-between">
              <span className="text-emerald-400 font-medium">{m.sender}</span>
              <Badge variant="secondary">{m.topic}</Badge>
            </div>
            <p className="text-gray-300">{JSON.stringify(m.payload, null, 2)}</p>
            <p className="text-xs text-gray-500">{new Date(m.timestamp).toLocaleTimeString()}</p>
          </div>
        )))
        }
      </CardContent>
    </Card>
  );
}