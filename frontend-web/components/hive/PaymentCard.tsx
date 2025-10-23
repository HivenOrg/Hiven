interface PaymentCardProps {
  payment: {
    id: number;
    title: string;
    paidBy: string;
    amount: number;
    icon: React.ElementType;
    iconBg: string;
  };
}

export default function PaymentCard({ payment }: PaymentCardProps) {
  const Icon = payment.icon;
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
    </div>
  );
}
