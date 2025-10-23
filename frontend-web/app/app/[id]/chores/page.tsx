"use client";

import { useState } from "react";
import ChoreCard from "@/components/hive/ChoreCard";
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

// Mock data for members - replace with actual API call
const members = [
  { id: 1, name: "Priya Patel" },
  { id: 2, name: "Rahul Sharma" },
  { id: 3, name: "Amit Kumar" },
  { id: 4, name: "Sneha Reddy" },
];

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
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [choreName, setChoreName] = useState("");
  const [assignedMember, setAssignedMember] = useState("");
  const [dueDate, setDueDate] = useState("");

  const handleAddChore = () => {
    // Implement add chore logic
    console.log("Adding chore:", {
      name: choreName,
      assignedMember,
      dueDate,
    });
    // Reset form
    setChoreName("");
    setAssignedMember("");
    setDueDate("");
    setIsModalOpen(false);
  };

  const isFormValid = choreName && assignedMember && dueDate;

  return (
    <div className="p-4">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Chores</h2>
        <Button size="sm" variant="default" onClick={() => setIsModalOpen(true)}>
          Add Chore
        </Button>
      </div>

      {/* Chores List */}
      <div className="space-y-3">
        {chores.map((chore) => (
          <ChoreCard key={chore.id} chore={chore} />
        ))}
      </div>

      {/* Add Chore Modal */}
      <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add Chore</DialogTitle>
            <DialogDescription>
              Create a new chore and assign it to a member.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            {/* Chore Name */}
            <div className="space-y-2">
              <Label htmlFor="chore-name">Chore Name</Label>
              <Input
                id="chore-name"
                placeholder="e.g., Clean the kitchen"
                value={choreName}
                onChange={(e) => setChoreName(e.target.value)}
                className="bg-card border-border"
              />
            </div>

            {/* Assign to Member */}
            <div className="space-y-2">
              <Label htmlFor="assigned-member">Assign To</Label>
              <Select value={assignedMember} onValueChange={setAssignedMember}>
                <SelectTrigger id="assigned-member" className="bg-card border-border">
                  <SelectValue placeholder="Select a member" />
                </SelectTrigger>
                <SelectContent>
                  {members.map((member) => (
                    <SelectItem key={member.id} value={member.id.toString()}>
                      {member.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>

            {/* Due Date */}
            <div className="space-y-2">
              <Label htmlFor="due-date">Due Date</Label>
              <Input
                id="due-date"
                type="date"
                value={dueDate}
                onChange={(e) => setDueDate(e.target.value)}
                className="bg-card border-border"
              />
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => {
                setChoreName("");
                setAssignedMember("");
                setDueDate("");
                setIsModalOpen(false);
              }}
            >
              Cancel
            </Button>
            <Button onClick={handleAddChore} disabled={!isFormValid}>
              Add Chore
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
