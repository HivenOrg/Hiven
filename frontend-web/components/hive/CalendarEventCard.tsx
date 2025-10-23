interface CalendarEventCardProps {
    event: {
        id: number;
        name: string;
        date: string;
        icon: React.ElementType;
    };
}

export default function CalendarEventCard({ event }: CalendarEventCardProps) {
  const Icon = event.icon;
  return (
    <div
      key={event.id}
      className="flex items-center gap-4 p-4 rounded-xl bg-card border border-border"
    >
      {/* Icon */}
      <div className="w-12 h-12 rounded-full bg-primary/20 flex items-center justify-center shrink-0">
        <Icon className="w-6 h-6 text-primary" />
      </div>

      {/* Event Info */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold text-base text-foreground">
          {event.name}
        </h3>
        <p className="text-sm text-muted-foreground">{event.date}</p>
      </div>
    </div>
  );
}
