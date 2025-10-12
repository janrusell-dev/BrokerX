export function formatTimestamp(timestamp: string): string{
    return new Date(timestamp).toLocaleTimeString();
}

export function formatUptime(seconds: number): string{
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    const secs = Math.floor(seconds % 60);

    if (hours > 0){
        return `${hours}h ${minutes}m ${secs}s`;
    }

    if (minutes > 0){
        return `${minutes}m ${secs}s`;
    }
    return `${secs}s`;
}

export function formatMessageRate(rate: number): string{
    if (rate >= 1000){
        return `${(rate / 1000).toFixed(1)}k/s`;
    }
    return `${rate.toFixed(1)}/s`
}