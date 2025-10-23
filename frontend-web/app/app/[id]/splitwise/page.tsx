"use client";

import { useState } from "react";
import { Home, Lightbulb, Pizza, ShoppingBag } from "lucide-react";
import OwedBalanceCard from "@/components/hive/OwedBalanceCard";
import PaymentCard from "@/components/hive/PaymentCard";
import TransactionCard from "@/components/hive/TransactionCard";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

// Mock data for members - replace with actual API call
const members = [
  { id: 1, name: "You" },
  { id: 2, name: "Priya Patel" },
  { id: 3, name: "Rahul Sharma" },
  { id: 4, name: "Amit Kumar" },
  { id: 5, name: "Sneha Reddy" },
];

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
  const [isPaymentModalOpen, setIsPaymentModalOpen] = useState(false);
  const [paymentTitle, setPaymentTitle] = useState("");
  const [paymentAmount, setPaymentAmount] = useState("");
  const [paidBy, setPaidBy] = useState("");

  const [isTransactionModalOpen, setIsTransactionModalOpen] = useState(false);
  const [transactionAmount, setTransactionAmount] = useState("");
  const [transactionMember, setTransactionMember] = useState("");
  const [transactionType, setTransactionType] = useState("sending");

  const handleAddPayment = () => {
    // Implement add payment logic
    console.log("Adding payment:", {
      title: paymentTitle,
      amount: parseFloat(paymentAmount),
      paidBy,
    });
    // Reset form
    setPaymentTitle("");
    setPaymentAmount("");
    setPaidBy("");
    setIsPaymentModalOpen(false);
  };

  const handleAddTransaction = () => {
    // Implement add transaction logic
    console.log("Adding transaction:", {
      amount: parseFloat(transactionAmount),
      member: transactionMember,
      type: transactionType,
      from: transactionType === "sending" ? "You" : transactionMember,
      to: transactionType === "sending" ? transactionMember : "You",
    });
    // Reset form
    setTransactionAmount("");
    setTransactionMember("");
    setTransactionType("sending");
    setIsTransactionModalOpen(false);
  };

  const isPaymentFormValid = paymentTitle && paymentAmount && paidBy;
  const isTransactionFormValid = transactionAmount && transactionMember;

  return (
    <div className="p-4 space-y-8">
      {/* Payments Section */}
      <section>
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-xl font-semibold">Payments</h2>
          <Button
            size="sm"
            variant="default"
            onClick={() => setIsPaymentModalOpen(true)}
          >
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
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-xl font-semibold">Transactions</h2>
          <Button
            size="sm"
            variant="default"
            onClick={() => setIsTransactionModalOpen(true)}
          >
            Add Transaction
          </Button>
        </div>
        <div className="space-y-3">
          {transactions.map((transaction) => (
            <TransactionCard key={transaction.id} transaction={transaction} />
          ))}
        </div>
      </section>

      {/* Add Payment Modal */}
      <Dialog open={isPaymentModalOpen} onOpenChange={setIsPaymentModalOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add Payment</DialogTitle>
            <DialogDescription>
              Record a payment made by a member of your hive.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            {/* Payment Title */}
            <div className="space-y-2">
              <Label htmlFor="payment-title">Payment Title</Label>
              <Input
                id="payment-title"
                placeholder="e.g., Groceries"
                value={paymentTitle}
                onChange={(e) => setPaymentTitle(e.target.value)}
                className="bg-card border-border"
              />
            </div>

            {/* Amount */}
            <div className="space-y-2">
              <Label htmlFor="payment-amount">Amount</Label>
              <Input
                id="payment-amount"
                type="number"
                step="0.01"
                placeholder="0.00"
                value={paymentAmount}
                onChange={(e) => setPaymentAmount(e.target.value)}
                className="bg-card border-border"
              />
            </div>

            {/* Paid By */}
            <div className="space-y-2">
              <Label htmlFor="paid-by">Paid By</Label>
              <Select value={paidBy} onValueChange={setPaidBy}>
                <SelectTrigger id="paid-by" className="bg-card border-border">
                  <SelectValue placeholder="Select a member" />
                </SelectTrigger>
                <SelectContent>
                  {members.map((member) => (
                    <SelectItem key={member.id} value={member.name}>
                      {member.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => {
                setPaymentTitle("");
                setPaymentAmount("");
                setPaidBy("");
                setIsPaymentModalOpen(false);
              }}
            >
              Cancel
            </Button>
            <Button onClick={handleAddPayment} disabled={!isPaymentFormValid}>
              Add Payment
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Add Transaction Modal */}
      <Dialog
        open={isTransactionModalOpen}
        onOpenChange={setIsTransactionModalOpen}
      >
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add Transaction</DialogTitle>
            <DialogDescription>
              Record a money transfer between you and another member to settle
              debts.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            {/* Transaction Type */}
            <div className="space-y-2">
              <Label htmlFor="transaction-type">Transaction Type</Label>
              <Select
                value={transactionType}
                onValueChange={setTransactionType}
              >
                <SelectTrigger
                  id="transaction-type"
                  className="bg-card border-border"
                >
                  <SelectValue placeholder="Select type" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="sending">Sending Money</SelectItem>
                  <SelectItem value="receiving">Receiving Money</SelectItem>
                </SelectContent>
              </Select>
            </div>

            {/* Member */}
            <div className="space-y-2">
              <Label htmlFor="transaction-member">
                {transactionType === "sending" ? "Send To" : "Receive From"}
              </Label>
              <Select
                value={transactionMember}
                onValueChange={setTransactionMember}
              >
                <SelectTrigger
                  id="transaction-member"
                  className="bg-card border-border"
                >
                  <SelectValue placeholder="Select a member" />
                </SelectTrigger>
                <SelectContent>
                  {members
                    .filter((member) => member.name !== "You")
                    .map((member) => (
                      <SelectItem key={member.id} value={member.name}>
                        {member.name}
                      </SelectItem>
                    ))}
                </SelectContent>
              </Select>
            </div>

            {/* Amount */}
            <div className="space-y-2">
              <Label htmlFor="transaction-amount">Amount</Label>
              <Input
                id="transaction-amount"
                type="number"
                step="0.01"
                placeholder="0.00"
                value={transactionAmount}
                onChange={(e) => setTransactionAmount(e.target.value)}
                className="bg-card border-border"
              />
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => {
                setTransactionAmount("");
                setTransactionMember("");
                setTransactionType("sending");
                setIsTransactionModalOpen(false);
              }}
            >
              Cancel
            </Button>
            <Button
              onClick={handleAddTransaction}
              disabled={!isTransactionFormValid}
            >
              Add Transaction
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
