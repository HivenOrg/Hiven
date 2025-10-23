import { MoreVertical } from "lucide-react";
import { Button } from "../ui/button";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "../ui/dropdown-menu";

interface ChoreCardProps {
  chore: {
    id: number;
    name: string;
    dueInDays: number;
    assignee: {
      id: number;
      name: string;
      profilePicture: string;
    };
  };
  onEdit?: (chore: ChoreCardProps["chore"]) => void;
  onDelete?: (choreId: number) => void;
  onMarkComplete?: (choreId: number) => void;
}

export default function ChoreCard({
  chore,
  onEdit,
  onDelete,
  onMarkComplete,
}: ChoreCardProps) {
  const handleMarkComplete = () => {
    if (onMarkComplete) {
      onMarkComplete(chore.id);
    }
  };

  const handleEdit = () => {
    if (onEdit) {
      onEdit(chore);
    }
  };

  const handleDelete = () => {
    if (onDelete) {
      onDelete(chore.id);
    }
  };

  return (
    <div className="flex items-center gap-4 p-4 rounded-xl bg-card border border-border">
      {/* Assignee Profile Picture */}
      <div className="relative w-12 h-12 rounded-full overflow-hidden bg-muted shrink-0">
        <img
          src={chore.assignee.profilePicture}
          alt={chore.assignee.name}
          className="object-cover"
          onError={(e) => {
            e.currentTarget.style.display = "none";
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-primary/10" />
      </div>

      {/* Chore Info */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold text-base text-foreground truncate">
          {chore.name}
        </h3>
        <p className="text-sm text-muted-foreground truncate">
          {chore.assignee.name}
        </p>
        <p className="text-sm text-primary font-medium">
          Due in {chore.dueInDays} days
        </p>
      </div>

      {/* Actions Menu */}
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button size="icon" variant="ghost" className="shrink-0 h-8 w-8">
            <MoreVertical className="w-5 h-5" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" className="w-48">
          <DropdownMenuItem onClick={handleMarkComplete}>
            Mark as Complete
          </DropdownMenuItem>
          <DropdownMenuItem onClick={handleEdit}>
            Edit
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={handleDelete}
            className="text-destructive focus:text-destructive"
          >
            Delete Chore
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
}
