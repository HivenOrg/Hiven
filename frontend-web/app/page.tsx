import Link from "next/link";
import {
  Home as HomeIcon,
  Users,
  ListChecks,
  DollarSign,
  Calendar,
  Sparkles,
  CheckCircle2,
  ArrowRight,
} from "lucide-react";
import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      {/* Hero Section */}
      <section className="relative overflow-hidden px-4 pt-16 pb-20">
        {/* Gradient Background Effects */}
        <div className="absolute top-0 left-1/2 -translate-x-1/2 w-[600px] h-[600px] bg-primary/10 rounded-full blur-[120px] -z-10" />
        <div className="absolute top-40 right-0 w-[400px] h-[400px] bg-primary/5 rounded-full blur-[100px] -z-10" />

        <div className="max-w-4xl mx-auto text-center">
          {/* Logo/Brand */}
          <div className="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-primary/10 border border-primary/20 mb-8">
            <Sparkles className="w-4 h-4 text-primary" />
            <span className="text-sm font-medium text-primary">
              Welcome to Hiven
            </span>
          </div>

          {/* Main Headline */}
          <h1 className="text-5xl md:text-7xl font-bold mb-6 leading-tight">
            Shared living,
            <br />
            <span className="text-primary">simplified</span>
          </h1>

          {/* Subtitle */}
          <p className="text-lg md:text-xl text-muted-foreground mb-12 max-w-2xl mx-auto leading-relaxed">
            Manage your shared space effortlessly. Track chores, split expenses,
            schedule events, and stay connected with your housemates—all in one
            place.
          </p>

          {/* CTA Buttons */}
          <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
            <Link href="/signup">
              <Button
                size="lg"
                className="rounded-full px-8 py-6 text-base font-semibold w-full sm:w-auto group"
              >
                Get Started
                <ArrowRight className="ml-2 w-4 h-4 group-hover:translate-x-1 transition-transform" />
              </Button>
            </Link>
            <Link href="/login">
              <Button
                size="lg"
                variant="outline"
                className="rounded-full px-8 py-6 text-base font-semibold w-full sm:w-auto border-border hover:bg-card"
              >
                Sign In
              </Button>
            </Link>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="px-4 py-20 bg-card/30">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-5xl font-bold mb-4">
              Everything you need for shared living
            </h2>
            <p className="text-muted-foreground text-lg">
              Powerful features designed to make household management effortless
            </p>
          </div>

          {/* Feature Grid */}
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            {/* Feature 1: Members */}
            <div className="bg-card rounded-2xl p-6 border border-border hover:border-primary/30 transition-all group">
              <div className="w-12 h-12 rounded-xl bg-primary/10 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <Users className="w-6 h-6 text-primary" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Member Management</h3>
              <p className="text-muted-foreground text-sm">
                Add, manage, and coordinate with all your housemates in one
                central hub.
              </p>
            </div>

            {/* Feature 2: Chores */}
            <div className="bg-card rounded-2xl p-6 border border-border hover:border-primary/30 transition-all group">
              <div className="w-12 h-12 rounded-xl bg-blue-500/10 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <ListChecks className="w-6 h-6 text-blue-500" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Chore Tracking</h3>
              <p className="text-muted-foreground text-sm">
                Assign tasks, set due dates, and ensure everyone does their fair
                share.
              </p>
            </div>

            {/* Feature 3: Splitwise */}
            <div className="bg-card rounded-2xl p-6 border border-border hover:border-primary/30 transition-all group">
              <div className="w-12 h-12 rounded-xl bg-green-500/10 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <DollarSign className="w-6 h-6 text-green-500" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Expense Splitting</h3>
              <p className="text-muted-foreground text-sm">
                Track payments, split bills fairly, and settle debts with
                ease.
              </p>
            </div>

            {/* Feature 4: Calendar */}
            <div className="bg-card rounded-2xl p-6 border border-border hover:border-primary/30 transition-all group">
              <div className="w-12 h-12 rounded-xl bg-purple-500/10 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                <Calendar className="w-6 h-6 text-purple-500" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Shared Calendar</h3>
              <p className="text-muted-foreground text-sm">
                Schedule events, set reminders, and keep everyone in sync.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* How It Works Section */}
      <section className="px-4 py-20">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-5xl font-bold mb-4">
              Get started in minutes
            </h2>
            <p className="text-muted-foreground text-lg">
              Set up your hive and start managing your shared space
            </p>
          </div>

          {/* Steps */}
          <div className="space-y-12">
            {/* Step 1 */}
            <div className="flex flex-col md:flex-row gap-6 items-start">
              <div className="flex-shrink-0">
                <div className="w-12 h-12 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                  1
                </div>
              </div>
              <div className="flex-1">
                <div className="flex items-center gap-3 mb-2">
                  <HomeIcon className="w-6 h-6 text-primary" />
                  <h3 className="text-2xl font-semibold">Create your Hive</h3>
                </div>
                <p className="text-muted-foreground">
                  Sign up and create a hive for your shared living space. Give
                  it a name and you're ready to go.
                </p>
              </div>
            </div>

            {/* Step 2 */}
            <div className="flex flex-col md:flex-row gap-6 items-start">
              <div className="flex-shrink-0">
                <div className="w-12 h-12 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                  2
                </div>
              </div>
              <div className="flex-1">
                <div className="flex items-center gap-3 mb-2">
                  <Users className="w-6 h-6 text-primary" />
                  <h3 className="text-2xl font-semibold">Invite members</h3>
                </div>
                <p className="text-muted-foreground">
                  Add your housemates by email. They'll get an invitation to
                  join your hive instantly.
                </p>
              </div>
            </div>

            {/* Step 3 */}
            <div className="flex flex-col md:flex-row gap-6 items-start">
              <div className="flex-shrink-0">
                <div className="w-12 h-12 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                  3
                </div>
              </div>
              <div className="flex-1">
                <div className="flex items-center gap-3 mb-2">
                  <CheckCircle2 className="w-6 h-6 text-primary" />
                  <h3 className="text-2xl font-semibold">
                    Start organizing
                  </h3>
                </div>
                <p className="text-muted-foreground">
                  Assign chores, track expenses, schedule events, and keep
                  everyone on the same page effortlessly.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="px-4 py-20 relative overflow-hidden">
        {/* Gradient Background */}
        <div className="absolute inset-0 bg-gradient-to-br from-primary/10 via-transparent to-primary/5 -z-10" />

        <div className="max-w-4xl mx-auto text-center">
          <h2 className="text-3xl md:text-5xl font-bold mb-6">
            Ready to simplify your shared living?
          </h2>
          <p className="text-muted-foreground text-lg mb-12 max-w-2xl mx-auto">
            Join hundreds of households already using Hiven to make shared
            living stress-free.
          </p>
          <Link href="/signup">
            <Button
              size="lg"
              className="rounded-full px-8 py-6 text-base font-semibold group"
            >
              Create Your Hive
              <ArrowRight className="ml-2 w-4 h-4 group-hover:translate-x-1 transition-transform" />
            </Button>
          </Link>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-border px-4 py-12">
        <div className="max-w-6xl mx-auto">
          <div className="flex flex-col md:flex-row justify-between items-center gap-6">
            <div className="flex items-center gap-2">
              <Sparkles className="w-6 h-6 text-primary" />
              <span className="text-xl font-bold">Hiven</span>
            </div>
            <p className="text-muted-foreground text-sm">
              © 2025 Hiven. All rights reserved.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
