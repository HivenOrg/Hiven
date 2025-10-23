"use client";

import { useParams, usePathname } from "next/navigation";
import HiveLayoutHeader from "@/components/hive/layout/header";
import HiveBottomNavigation from "@/components/hive/layout/BottomNavigation";

interface HivePageLayoutProps {
  children: React.ReactNode;
}

// Mock hive data - replace with actual API call
const getHiveName = (id: string) => {
  const hives: Record<string, string> = {
    "1": "Mumbai Flat",
    "2": "Delhi Home",
  };
  return hives[id] || "Hive";
};

export default function HivePageLayout({ children }: HivePageLayoutProps) {
  const params = useParams();
  const pathname = usePathname();
  const hiveId = params.id as string;
  const hiveName = getHiveName(hiveId);

  return (
    <div className="hive-page-layout flex flex-col h-screen bg-background">
      {/* Header */}
      <HiveLayoutHeader hiveName={hiveName} />

      {/* Main Content */}
      <main className="flex-1 overflow-y-auto pb-20">{children}</main>

      {/* Bottom Navigation */}
      <HiveBottomNavigation hiveId={hiveId} pathname={pathname} />
    </div>
  );
}
