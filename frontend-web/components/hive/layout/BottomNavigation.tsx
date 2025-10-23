import { Calendar, Home, ListChecks, Wallet } from "lucide-react";
import Link from "next/link";
import { cn } from "@/lib/utils";

interface HiveBottomNavigationProps {
  hiveId: string;
  pathname: string;
}

export default function HiveBottomNavigation({
  hiveId,
  pathname,
}: HiveBottomNavigationProps) {
  // Navigation items
  const navItems = [
    {
      name: "Home",
      href: `/app/${hiveId}`,
      icon: Home,
      isActive: pathname === `/app/${hiveId}`,
    },
    {
      name: "Chores",
      href: `/app/${hiveId}/chores`,
      icon: ListChecks,
      isActive: pathname === `/app/${hiveId}/chores`,
    },
    {
      name: "Splitwise",
      href: `/app/${hiveId}/splitwise`,
      icon: Wallet,
      isActive: pathname === `/app/${hiveId}/splitwise`,
    },
    {
      name: "Calendar",
      href: `/app/${hiveId}/calendar`,
      icon: Calendar,
      isActive: pathname === `/app/${hiveId}/calendar`,
    },
  ];
  return (
    <nav className="fixed bottom-0 left-0 right-0 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 border-t border-border safe-area-inset-bottom">
      <div className="max-w-lg mx-auto px-2 py-2">
        <div className="flex items-center justify-around gap-1">
          {navItems.map((item) => {
            const Icon = item.icon;
            return (
              <Link
                key={item.name}
                href={item.href}
                className={cn(
                  "flex flex-col items-center gap-1 px-4 py-2 rounded-lg transition-all active:scale-95",
                  item.isActive
                    ? "text-primary bg-primary/10 rounded-full"
                    : "text-muted-foreground hover:text-foreground hover:bg-accent/50",
                )}
              >
                <Icon
                  className={cn(
                    "w-5 h-5",
                    item.isActive ? "stroke-[2.5]" : "stroke-2",
                  )}
                />
                <span className="text-xs font-medium">{item.name}</span>
              </Link>
            );
          })}
        </div>
      </div>
    </nav>
  );
}
