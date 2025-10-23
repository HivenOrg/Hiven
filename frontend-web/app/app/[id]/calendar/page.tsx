"use client";

import { Calendar, PartyPopper, Sparkles, Utensils, Wallet } from "lucide-react";
import { Button } from "@/components/ui/button";

// Mock data for events - replace with actual API call later
const events = [
  {
    id: 1,
    name: "Parent Visit",
    date: "Oct 24",
    icon: Calendar,
  },
  {
    id: 2,
    name: "Rent Due",
    date: "Oct 25",
    icon: Wallet,
  },
  {
    id: 3,
    name: "Cleaning Day",
    date: "Oct 26",
    icon: Sparkles,
  },
  {
    id: 4,
    name: "Movie Night",
    date: "Oct 28",
    icon: PartyPopper,
  },
  {
    id: 5,
    name: "Grocery Shopping",
    date: "Oct 30",
    icon: Utensils,
  },
];

export default function CalendarPage() {
  return (
    <div className="p-4">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Calendar</h2>
        <Button size="sm" variant="default">
          Add Event
        </Button>
      </div>

      {/* Events List */}
      <div className="space-y-3">
        {events.map((event) => {
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
        })}
      </div>
    </div>
  );
}
