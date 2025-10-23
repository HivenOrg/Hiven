'use client'

import HiveCard from "@/components/app/hiveCard";
import { Button } from "@/components/ui/button";
import { Plus, ChevronRight } from "lucide-react";
import Image from "next/image";

// Mock data for hives - replace with actual API call later
const hives = [
  {
    id: 1,
    name: "Mumbai Flat",
    memberCount: 4,
    image: "/hive-images/mumbai-flat.jpg", // Placeholder - replace with actual images
  },
  {
    id: 2,
    name: "Delhi Home",
    memberCount: 3,
    image: "/hive-images/delhi-home.jpg", // Placeholder - replace with actual images
  },
];

export default function AppPage() {
  return (
    <div className="min-h-screen bg-background">
      {/* Navbar */}
      <header className="sticky top-0 z-10 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 border-b border-border">
        <div className="flex items-center justify-between px-4 py-3">
          {/* Profile Picture */}
          <div className="relative w-10 h-10 rounded-full overflow-hidden bg-muted">
            <img
              src="https://placehold.co/100" // Replace with actual profile image
              alt="Profile"
              className="object-cover"
              onError={(e) => {
                // Fallback to a colored circle if image fails
                e.currentTarget.style.display = "none";
              }}
            />
            <div className="absolute inset-0 bg-gradient-to-br from-primary/40 to-primary/20" />
          </div>

          {/* Title */}
          <h1 className="text-lg font-semibold">My Hives</h1>

          {/* Add Button */}
          <Button size="icon" variant="ghost" className="rounded-full active:scale-[0.9] active:bg-card transition-all">
            <Plus className="w-6 h-6" />
          </Button>
        </div>
      </header>

      {/* Hives List */}
      <main className="p-4 space-y-4">
        {hives.map((hive) => (
          <HiveCard key={hive.id} hive={hive} />
        ))}
      </main>
    </div>
  );
}