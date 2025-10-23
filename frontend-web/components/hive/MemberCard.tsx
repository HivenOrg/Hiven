import { Crown, X } from "lucide-react";
import { Button } from "@/components/ui/button";

interface MemberCardProps {
  member: {
    id: number;
    name: string;
    email: string;
    profilePicture: string;
  };
  owner?: boolean;
  onRemove?: (memberId: number) => void;
  canRemove?: boolean;
}

export default function MemberCard({
  member,
  owner,
  onRemove,
  canRemove,
}: MemberCardProps) {
  return (
    <div className="flex items-center gap-4 p-4 rounded-xl bg-card border border-border transition-all hover:bg-card/80">
      {/* Profile Picture */}
      <div className="relative w-12 h-12 rounded-full overflow-hidden bg-muted shrink-0">
        <img
          src={member.profilePicture}
          alt={member.name}
          className="object-cover"
          onError={(e) => {
            e.currentTarget.style.display = "none";
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-primary/10" />
      </div>

      {/* Member Info */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold flex items-center gap-2 text-base text-foreground truncate">
          {member.name} {owner && <Crown className="h-4 w-4 text-primary" />}
        </h3>
        <p className="text-sm text-muted-foreground truncate">{member.email}</p>
      </div>

      {/* Remove Button */}
      {canRemove && onRemove && !owner && (
        <Button
          size="icon"
          variant="ghost"
          className="shrink-0 h-8 w-8 text-destructive hover:text-destructive hover:bg-destructive/10"
          onClick={() => onRemove(member.id)}
        >
          <X className="w-5 h-5" />
        </Button>
      )}
    </div>
  );
}
