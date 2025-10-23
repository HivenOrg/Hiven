"use client";

import {
  Calendar,
  PartyPopper,
  Sparkles,
  Utensils,
  Wallet,
} from "lucide-react";
import CalendarEventCard from "@/components/hive/CalendarEventCard";
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
          return <CalendarEventCard key={event.id} event={event} />;
        })}
      </div>
    </div>
  );
}
