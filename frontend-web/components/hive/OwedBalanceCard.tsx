
interface OwedBalanceCardProps{
    balance: {
        id: number;
        person: {
            name: string;
            profilePicture: string;
        };
        amount: number;
        owesYou: boolean;
        reason: string;
    };
} 

export default function OwedBalanceCard({ balance }: OwedBalanceCardProps) {
  return (
    <div
      key={balance.id}
      className="flex items-center gap-4 p-4 rounded-xl bg-card border border-border"
    >
      {/* Profile Picture */}
      <div className="relative w-12 h-12 rounded-full overflow-hidden bg-muted shrink-0">
        <img
          src={balance.person.profilePicture}
          alt={balance.person.name}
          className="object-cover"
          onError={(e) => {
            e.currentTarget.style.display = "none";
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-primary/10" />
      </div>

      {/* Balance Info */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold text-base text-foreground">
          {balance.owesYou
            ? `${balance.person.name} owes you`
            : `You owe ${balance.person.name}`}
        </h3>
        <p className="text-sm text-muted-foreground">{balance.reason}</p>
      </div>

      {/* Amount */}
      <div
        className={`text-lg font-semibold ${
          balance.owesYou ? "text-green-500" : "text-red-500"
        }`}
      >
        {balance.owesYou ? "+" : "-"}${balance.amount.toFixed(2)}
      </div>
    </div>
  );
}
