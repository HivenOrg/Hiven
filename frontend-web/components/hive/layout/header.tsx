import { ChevronLeft } from "lucide-react";
import Link from "next/link";
import { Button } from "@/components/ui/button";

interface HiveLayoutHeaderProps {
  hiveName?: string;
}

export default function HiveLayoutHeader({ hiveName }: HiveLayoutHeaderProps) {
  return (
    <header className="sticky top-0 z-10 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 border-b border-border">
      <div className="flex items-center gap-3 px-4 py-3">
        {/* Back Button */}
        <Link href="/app">
          <Button
            size="icon"
            variant="ghost"
            className="rounded-full active:scale-[0.9] transition-all"
          >
            <ChevronLeft className="w-6 h-6" />
          </Button>
        </Link>

        {/* Hive Name */}
        <h1 className="flex-1 text-lg font-semibold text-center pr-10">
          {hiveName}
        </h1>
      </div>
    </header>
  );
}
