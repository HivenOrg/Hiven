import { MoreVertical } from "lucide-react";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";

interface PaymentCardProps {
  payment: {
    id: number;
    title: string;
    paidBy: string;
    amount: number;
    icon: React.ElementType;
    iconBg: string;
  };
  onEdit?: (payment: PaymentCardProps["payment"]) => void;
  onDelete?: (paymentId: number) => void;
}

export default function PaymentCard({
  payment,
  onEdit,
  onDelete,
}: PaymentCardProps) {
  const Icon = payment.icon;

  const handleEdit = () => {
    if (onEdit) {
      onEdit(payment);
    }
  };

  const handleDelete = () => {
    if (onDelete) {
      onDelete(payment.id);
    }
  };

  return (
    <div
      key={payment.id}
      className="flex items-center gap-4 p-4 rounded-xl bg-card border border-border"
    >
      {/* Icon */}
      <div
        className={`w-12 h-12 rounded-full flex items-center justify-center shrink-0 ${payment.iconBg}`}
      >
        <Icon className="w-6 h-6 text-foreground" />
      </div>

      {/* Payment Info */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold text-base text-foreground">
          {payment.title}
        </h3>
        <p className="text-sm text-muted-foreground">{payment.paidBy} paid</p>
      </div>

      {/* Amount */}
      <div className="text-lg font-semibold text-foreground">
        ${payment.amount.toFixed(2)}
      </div>

      {/* Actions Menu */}
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button size="icon" variant="ghost" className="shrink-0 h-8 w-8">
            <MoreVertical className="w-5 h-5" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" className="w-48">
          <DropdownMenuItem onClick={handleEdit}>Edit</DropdownMenuItem>
          <DropdownMenuItem
            onClick={handleDelete}
            className="text-destructive focus:text-destructive"
          >
            Delete
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
}
