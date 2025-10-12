export interface PublishRequest {
    topic: string;
    sender: string;
    payload: Record<string, any>;
}

export interface PublishResponse {
    status: string;
    message: string;
    topic: string;
    latency: number;
}

export interface MetricsResponse {
    totalMessages: number;
    activeSubscribers: number;
    avgLatency:number;
    messageRate: number;
    uptime: number;
    topicMetrics: Record<string, TopicMetric>;
    latencyHistory: LatencyPoint[];
}

export interface TopicMetric{
    messageCount: number;
    avgLatency: number;
}

export interface LatencyPoint{
    timestamp: string;
    latency: number;
}

export interface Message{
    topic: string;
    sender: string;
    payload: Record<string, any>;
    timestamp: string;
}

export interface TopicInfo {
    exists : boolean;
    subscribers: number;
    messageCount: number;
    lastPublished: string | null;
}