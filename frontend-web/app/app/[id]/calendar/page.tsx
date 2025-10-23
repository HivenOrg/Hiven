"use client";

import { useState } from "react";
import {
  Calendar,
  PartyPopper,
  Sparkles,
  Utensils,
  Wallet,
} from "lucide-react";
import CalendarEventCard from "@/components/hive/CalendarEventCard";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

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
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [eventName, setEventName] = useState("");
  const [eventDate, setEventDate] = useState("");

  const handleAddEvent = () => {
    // Implement add event logic
    console.log("Adding event:", {
      name: eventName,
      date: eventDate,
    });
    // Reset form
    setEventName("");
    setEventDate("");
    setIsModalOpen(false);
  };

  const isFormValid = eventName && eventDate;

  return (
    <div className="p-4">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Calendar</h2>
        <Button size="sm" variant="default" onClick={() => setIsModalOpen(true)}>
          Add Event
        </Button>
      </div>

      {/* Events List */}
      <div className="space-y-3">
        {events.map((event) => {
          return <CalendarEventCard key={event.id} event={event} />;
        })}
      </div>

      {/* Add Event Modal */}
      <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add Event</DialogTitle>
            <DialogDescription>
              Create a new event for your hive calendar.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            {/* Event Name */}
            <div className="space-y-2">
              <Label htmlFor="event-name">Event Name</Label>
              <Input
                id="event-name"
                placeholder="e.g., Movie Night"
                value={eventName}
                onChange={(e) => setEventName(e.target.value)}
                className="bg-card border-border"
              />
            </div>

            {/* Event Date */}
            <div className="space-y-2">
              <Label htmlFor="event-date">Date</Label>
              <Input
                id="event-date"
                type="date"
                value={eventDate}
                onChange={(e) => setEventDate(e.target.value)}
                className="bg-card border-border"
              />
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => {
                setEventName("");
                setEventDate("");
                setIsModalOpen(false);
              }}
            >
              Cancel
            </Button>
            <Button onClick={handleAddEvent} disabled={!isFormValid}>
              Add Event
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
