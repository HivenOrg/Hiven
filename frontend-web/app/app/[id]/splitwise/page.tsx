"use client";

import { ArrowRight, Home, Lightbulb, Pizza, ShoppingBag } from "lucide-react";
import { Button } from "@/components/ui/button";

// Mock data - replace with actual API call later
const payments = [
  {
    id: 1,
    title: "Pizza",
    paidBy: "You",
    amount: 12.5,
    icon: Pizza,
    iconBg: "bg-orange-500/20",
  },
  {
    id: 2,
    title: "Groceries",
    paidBy: "You",
    amount: 35.0,
    icon: ShoppingBag,
    iconBg: "bg-amber-500/20",
  },
  {
    id: 3,
    title: "Electricity",
    paidBy: "You",
    amount: 20.0,
    icon: Lightbulb,
    iconBg: "bg-yellow-500/20",
  },
  {
    id: 4,
    title: "Rent",
    paidBy: "You",
    amount: 1500.0,
    icon: Home,
    iconBg: "bg-orange-600/20",
  },
];

const owedBalances = [
  {
    id: 1,
    person: {
      name: "Liam",
      profilePicture: "https://placehold.co/100",
    },
    owesYou: true,
    amount: 5.0,
    reason: "for Pizza",
  },
  {
    id: 2,
    person: {
      name: "Sophia",
      profilePicture: "https://placehold.co/100",
    },
    owesYou: false,
    amount: 17.5,
    reason: "for Groceries",
  },
];

const transactions = [
  {
    id: 1,
    from: {
      name: "Liam",
      profilePicture: "https://placehold.co/100",
    },
    to: {
      name: "You",
      profilePicture: "https://placehold.co/100",
    },
    amount: 5.0,
    date: "Oct 20, 2025",
  },
  {
    id: 2,
    from: {
      name: "You",
      profilePicture: "https://placehold.co/100",
    },
    to: {
      name: "Amit",
      profilePicture: "https://placehold.co/100",
    },
    amount: 25.0,
    date: "Oct 18, 2025",
  },
];

export default function SplitwisePage() {
  return (
    <div className="p-4 space-y-8">
      {/* Payments Section */}
      <section>
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-xl font-semibold">Payments</h2>
          <Button size="sm" variant="default">
            Add Payment
          </Button>
        </div>
        <div className="space-y-3">
          {payments.map((payment) => {
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
                  <p className="text-sm text-muted-foreground">
                    {payment.paidBy} paid
                  </p>
                </div>

                {/* Amount */}
                <div className="text-lg font-semibold text-foreground">
                  ${payment.amount.toFixed(2)}
                </div>
              </div>
            );
          })}
        </div>
      </section>

      {/* Who Owes Who Section */}
      <section>
        <h2 className="text-xl font-semibold mb-4">Who Owes Who</h2>
        <div className="space-y-3">
          {owedBalances.map((balance) => (
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
                <p className="text-sm text-muted-foreground">
                  {balance.reason}
                </p>
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
          ))}
        </div>
      </section>

      {/* Transactions Section */}
      <section>
        <h2 className="text-xl font-semibold mb-4">Transactions</h2>
        <div className="space-y-3">
          {transactions.map((transaction) => (
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
                <p className="text-xs text-muted-foreground">
                  {transaction.date}
                </p>
              </div>

              {/* Amount */}
              <div className="text-base font-semibold text-foreground">
                ${transaction.amount.toFixed(2)}
              </div>
            </div>
          ))}
        </div>
      </section>
    </div>
  );
}
