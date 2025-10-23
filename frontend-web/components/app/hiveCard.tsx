import { ChevronRight } from "lucide-react";
import Image from "next/image";

interface HiveCardProps {
  hive: {
    id: number;
    name: string;
    memberCount: number;
    image: string;
  };
}

export default function HiveCard({ hive }: HiveCardProps) {
  return (
    <button
      type="button"
      className="w-full group relative overflow-hidden rounded-2xl bg-card border border-border p-4 flex items-center gap-4 transition-all hover:bg-card/80 hover:scale-[0.99] active:scale-[0.97]"
    >
      {/* Hive Image */}
      <div className="relative w-16 h-16 rounded-xl overflow-hidden bg-muted shrink-0">
        <Image
          src={hive.image}
          alt={hive.name}
          fill
          className="object-cover"
          onError={(e) => {
            // Fallback gradient if image fails to load
            e.currentTarget.style.display = "none";
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-br from-muted-foreground/10 to-muted-foreground/5" />
      </div>

      {/* Hive Info */}
      <div className="flex-1 text-left">
        <h2 className="font-semibold text-base text-foreground">{hive.name}</h2>
        <p className="text-sm text-primary font-medium">
          {hive.memberCount} members
        </p>
      </div>

      {/* Arrow Icon */}
      <ChevronRight className="w-5 h-5 text-muted-foreground group-hover:text-foreground transition-colors shrink-0" />
    </button>
  );
}
