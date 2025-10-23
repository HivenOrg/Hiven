"use client";

import { Camera, ChevronLeft, Lock, Mail, Phone, User } from "lucide-react";
import Link from "next/link";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

// Mock user data - replace with actual API call
const initialUserData = {
  name: "Rahul Sharma",
  email: "rahul.sharma@email.com",
  phone: "+91 98765 43210",
  profilePicture: "https://placehold.co/200",
};

export default function ProfilePage() {
  const [isEditing, setIsEditing] = useState(false);
  const [userData, setUserData] = useState(initialUserData);
  const [showPasswordChange, setShowPasswordChange] = useState(false);

  const handleSave = () => {
    // Implement save logic
    console.log("Saving user data:", userData);
    setIsEditing(false);
  };

  const handleCancel = () => {
    setUserData(initialUserData);
    setIsEditing(false);
    setShowPasswordChange(false);
  };

  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setUserData({ ...userData, profilePicture: reader.result as string });
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <header className="sticky top-0 z-10 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 border-b border-border">
        <div className="flex items-center gap-3 px-4 py-3">
          <Link href="/app">
            <Button
              size="icon"
              variant="ghost"
              className="rounded-full active:scale-[0.9] transition-all"
            >
              <ChevronLeft className="w-6 h-6" />
            </Button>
          </Link>
          <h1 className="flex-1 text-lg font-semibold text-center pr-10">
            Profile
          </h1>
        </div>
      </header>

      {/* Profile Content */}
      <main className="p-4 space-y-6">
        {/* Profile Picture Section */}
        <div className="flex flex-col items-center gap-4">
          <div className="relative">
            <div className="relative w-24 h-24 rounded-full overflow-hidden bg-muted">
              <img
                src={userData.profilePicture}
                alt={userData.name}
                className="object-cover w-full h-full"
                onError={(e) => {
                  e.currentTarget.style.display = "none";
                }}
              />
              <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-primary/10" />
            </div>
            {isEditing && (
              <>
                <input
                  type="file"
                  id="profile-picture-upload"
                  accept="image/*"
                  className="hidden"
                  onChange={handleImageUpload}
                />
                <label
                  htmlFor="profile-picture-upload"
                  className="absolute bottom-0 right-0 w-8 h-8 rounded-full bg-primary text-primary-foreground flex items-center justify-center shadow-lg hover:bg-primary/90 transition-all active:scale-95 cursor-pointer"
                >
                  <Camera className="w-4 h-4" />
                </label>
              </>
            )}
          </div>
          {!isEditing && (
            <Button
              variant="outline"
              size="sm"
              onClick={() => setIsEditing(true)}
            >
              Edit Profile
            </Button>
          )}
        </div>

        {/* Profile Form */}
        <div className="space-y-4">
          {/* Name */}
          <div className="space-y-2">
            <Label htmlFor="name" className="flex items-center gap-2">
              <User className="w-4 h-4 text-muted-foreground" />
              <span>Name</span>
            </Label>
            <Input
              id="name"
              value={userData.name}
              onChange={(e) =>
                setUserData({ ...userData, name: e.target.value })
              }
              disabled={!isEditing}
              className="bg-card border-border"
            />
          </div>

          {/* Email */}
          <div className="space-y-2">
            <Label htmlFor="email" className="flex items-center gap-2">
              <Mail className="w-4 h-4 text-muted-foreground" />
              <span>Email</span>
            </Label>
            <Input
              id="email"
              type="email"
              value={userData.email}
              onChange={(e) =>
                setUserData({ ...userData, email: e.target.value })
              }
              disabled={!isEditing}
              className="bg-card border-border"
            />
          </div>

          {/* Phone */}
          <div className="space-y-2">
            <Label htmlFor="phone" className="flex items-center gap-2">
              <Phone className="w-4 h-4 text-muted-foreground" />
              <span>Phone</span>
            </Label>
            <Input
              id="phone"
              type="tel"
              value={userData.phone}
              onChange={(e) =>
                setUserData({ ...userData, phone: e.target.value })
              }
              disabled={!isEditing}
              className="bg-card border-border"
            />
          </div>

          {/* Password Section */}
          {isEditing && (
            <div className="space-y-2">
              <Label className="flex items-center gap-2">
                <Lock className="w-4 h-4 text-muted-foreground" />
                <span>Password</span>
              </Label>
              {!showPasswordChange ? (
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setShowPasswordChange(true)}
                  className="w-full"
                >
                  Change Password
                </Button>
              ) : (
                <div className="space-y-3">
                  <Input
                    type="password"
                    placeholder="Current Password"
                    className="bg-card border-border"
                  />
                  <Input
                    type="password"
                    placeholder="New Password"
                    className="bg-card border-border"
                  />
                  <Input
                    type="password"
                    placeholder="Confirm New Password"
                    className="bg-card border-border"
                  />
                </div>
              )}
            </div>
          )}
        </div>

        {/* Action Buttons */}
        {isEditing && (
          <div className="flex gap-3 pt-4">
            <Button
              variant="outline"
              className="flex-1"
              onClick={handleCancel}
            >
              Cancel
            </Button>
            <Button className="flex-1" onClick={handleSave}>
              Save Changes
            </Button>
          </div>
        )}

        {/* Logout Button */}
        {!isEditing && (
          <div className="pt-8">
            <Button variant="destructive" className="w-full">
              Logout
            </Button>
          </div>
        )}
      </main>
    </div>
  );
}
