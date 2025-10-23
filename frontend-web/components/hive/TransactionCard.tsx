import { ArrowRight } from "lucide-react";

interface TransactionCardProps {
  transaction: {
    id: number;
    from: {
      name: string;
      profilePicture: string;
    };
    to: {
      name: string;
      profilePicture: string;
    };
    amount: number;
    date: string;
  };
}

export default function TransactionCard({ transaction }: TransactionCardProps) {
  return (
    <div
      key={transaction.id}
      className="flex items-center gap-3 p-4 rounded-xl bg-card border border-border"
    >
      {/* From Profile Picture */}
      <div className="relative w-10 h-10 rounded-full overflow-hidden bg-muted shrink-0">
        <img
          src={transaction.from.profilePicture}
          alt={transaction.from.name}
          className="object-cover"
          onError={(e) => {
            e.currentTarget.style.display = "none";
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-primary/10" />
      </div>

      {/* Arrow */}
      <ArrowRight className="w-5 h-5 text-muted-foreground shrink-0" />

      {/* To Profile Picture */}
      <div className="relative w-10 h-10 rounded-full overflow-hidden bg-muted shrink-0">
        <img
          src={transaction.to.profilePicture}
          alt={transaction.to.name}
          className="object-cover"
          onError={(e) => {
            e.currentTarget.style.display = "none";
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-primary/10" />
      </div>

      {/* Transaction Info */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold text-sm text-foreground">
          {transaction.from.name} → {transaction.to.name}
        </h3>
        <p className="text-xs text-muted-foreground">{transaction.date}</p>
      </div>

      {/* Amount */}
      <div className="text-base font-semibold text-foreground">
        ${transaction.amount.toFixed(2)}
      </div>
    </div>
  );
}
