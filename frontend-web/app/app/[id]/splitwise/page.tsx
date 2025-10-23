"use client";

import { Home, Lightbulb, Pizza, ShoppingBag } from "lucide-react";
import OwedBalanceCard from "@/components/hive/OwedBalanceCard";
import PaymentCard from "@/components/hive/PaymentCard";
import TransactionCard from "@/components/hive/TransactionCard";
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
            return <PaymentCard key={payment.id} payment={payment} />;
          })}
        </div>
      </section>

      {/* Who Owes Who Section */}
      <section>
        <h2 className="text-xl font-semibold mb-4">Who Owes Who</h2>
        <div className="space-y-3">
          {owedBalances.map((balance) => (
            <OwedBalanceCard key={balance.id} balance={balance} />
          ))}
        </div>
      </section>

      {/* Transactions Section */}
      <section>
        <h2 className="text-xl font-semibold mb-4">Transactions</h2>
        <div className="space-y-3">
          {transactions.map((transaction) => (
            <TransactionCard key={transaction.id} transaction={transaction} />
          ))}
        </div>
      </section>
    </div>
  );
}
