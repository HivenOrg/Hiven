"use client";

import ChoreCard from "@/components/hive/ChoreCard";
import { Button } from "@/components/ui/button";

// Mock data for chores - replace with actual API call later
const chores = [
  {
    id: 1,
    name: "Clean the kitchen",
    assignee: {
      id: 1,
      name: "Priya Patel",
      profilePicture: "https://placehold.co/100",
    },
    dueInDays: 2,
  },
  {
    id: 2,
    name: "Take out the trash",
    assignee: {
      id: 2,
      name: "Rahul Sharma",
      profilePicture: "https://placehold.co/100",
    },
    dueInDays: 3,
  },
  {
    id: 3,
    name: "Vacuum the living room",
    assignee: {
      id: 3,
      name: "Amit Kumar",
      profilePicture: "https://placehold.co/100",
    },
    dueInDays: 4,
  },
  {
    id: 4,
    name: "Clean the bathroom",
    assignee: {
      id: 4,
      name: "Sneha Reddy",
      profilePicture: "https://placehold.co/100",
    },
    dueInDays: 5,
  },
];

export default function ChoresPage() {

  return (
    <div className="p-4">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Chores</h2>
        <Button size="sm" variant="default">
          Add Chore
        </Button>
      </div>

      {/* Chores List */}
      <div className="space-y-3">
        {chores.map((chore) => (
          <ChoreCard key={chore.id} chore={chore} />
        ))}
      </div>
    </div>
  );
}
