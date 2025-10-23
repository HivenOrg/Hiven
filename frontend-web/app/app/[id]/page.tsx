"use client";

import { useState } from "react";
import { Plus } from "lucide-react";
import MemberCard from "@/components/hive/MemberCard";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

// Mock data for members - replace with actual API call later
const members = [
  {
    id: 1,
    name: "Rahul Sharma",
    email: "rahul.sharma@email.com",
    profilePicture: "https://placehold.co/100",
    owner: true,
  },
  {
    id: 2,
    name: "Priya Patel",
    email: "priya.patel@email.com",
    profilePicture: "https://placehold.co/100",
  },
  {
    id: 3,
    name: "Amit Kumar",
    email: "amit.kumar@email.com",
    profilePicture: "https://placehold.co/100",
  },
  {
    id: 4,
    name: "Sneha Reddy",
    email: "sneha.reddy@email.com",
    profilePicture: "https://placehold.co/100",
  },
];

export default function HivePage() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [memberEmail, setMemberEmail] = useState("");
  const [isRemoveDialogOpen, setIsRemoveDialogOpen] = useState(false);
  const [memberToRemove, setMemberToRemove] = useState<{
    id: number;
    name: string;
  } | null>(null);

  const handleAddMember = () => {
    // Implement add member logic
    console.log("Adding member with email:", memberEmail);
    setMemberEmail("");
    setIsModalOpen(false);
  };

  const handleRemoveClick = (memberId: number) => {
    const member = members.find((m) => m.id === memberId);
    if (member) {
      setMemberToRemove({ id: member.id, name: member.name });
      setIsRemoveDialogOpen(true);
    }
  };

  const handleConfirmRemove = () => {
    if (memberToRemove) {
      // Implement remove member logic
      console.log("Removing member with id:", memberToRemove.id);
      setMemberToRemove(null);
      setIsRemoveDialogOpen(false);
    }
  };

  const handleCancelRemove = () => {
    setMemberToRemove(null);
    setIsRemoveDialogOpen(false);
  };

  return (
    <div className="p-4 pb-24">
      {/* <h2 className="text-xl font-semibold mb-4">Members</h2> */}

      {/* Members List */}
      <div className="space-y-3">
        {members.map((member) => (
          <MemberCard
            key={member.id}
            member={member}
            owner={member.owner}
            onRemove={handleRemoveClick}
            canRemove={true}
          />
        ))}
      </div>

      {/* Floating Action Button */}
      <Button
        size="icon"
        className="fixed bottom-24 right-6 h-14 w-14 rounded-full shadow-lg hover:shadow-xl transition-all"
        onClick={() => setIsModalOpen(true)}
      >
        <Plus className="h-6 w-6" />
      </Button>

      {/* Add Member Modal */}
      <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add Member</DialogTitle>
            <DialogDescription>
              Enter the email address of the person you want to add to this
              hive.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div className="space-y-2">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                type="email"
                placeholder="member@email.com"
                value={memberEmail}
                onChange={(e) => setMemberEmail(e.target.value)}
                className="bg-card border-border"
              />
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => {
                setMemberEmail("");
                setIsModalOpen(false);
              }}
            >
              Cancel
            </Button>
            <Button onClick={handleAddMember} disabled={!memberEmail}>
              Add Member
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Remove Member Confirmation Dialog */}
      <AlertDialog open={isRemoveDialogOpen} onOpenChange={setIsRemoveDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Remove Member</AlertDialogTitle>
            <AlertDialogDescription>
              Are you sure you want to remove{" "}
              <span className="font-semibold text-foreground">
                {memberToRemove?.name}
              </span>{" "}
              from this hive? This action cannot be undone.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel onClick={handleCancelRemove}>
              Cancel
            </AlertDialogCancel>
            <AlertDialogAction
              onClick={handleConfirmRemove}
              className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
            >
              Remove
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
}
