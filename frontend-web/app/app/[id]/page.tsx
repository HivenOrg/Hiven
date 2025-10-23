"use client";

import MemberCard from "@/components/hive/MemberCard";

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
  return (
    <div className="p-4">
      {/* <h2 className="text-xl font-semibold mb-4">Members</h2> */}

      {/* Members List */}
      <div className="space-y-3">
        {members.map((member) => (
          <MemberCard key={member.id} member={member} owner={member.owner} />
        ))}
      </div>
    </div>
  );
}
