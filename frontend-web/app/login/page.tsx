import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import type { Metadata } from "next";
import Link from "next/link";

export const metadata: Metadata = {
  title: "Login",
  description: "Login to your account",
};

export default function LoginPage() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center  bg-background">
      <div className="w-full max-w-lg p-5">
        <h2 className="mb-6 text-center text-4xl font-bold text-primary">
          Welcome back
        </h2>
        <form className="space-y-6">
          <div>
            <Label htmlFor="email" className="mb-2 block text-sm font-medium">
              Email
            </Label>
            <Input id="email" placeholder="email" className="rounded-lg"/>
          </div>
          <div className="mb-1">
            <Label htmlFor="password" className="mb-2 block text-sm font-medium">
              Password
            </Label>
            <Input id="password" placeholder="password" type="password" className="rounded-lg" />
          </div>
          <div className="flex items-center justify-end">
            <Button variant="link" className="text-sm">
              Forgot password?
            </Button>
          </div>
          <Button
            type="submit"
            className="w-full rounded-full"
          >
            Login
          </Button>
          <Link href="/signup" className="text-sm text-muted-foreground hover:text-foreground text-center w-full block mt-2">
            Don't have an account? Sign up
          </Link>
        </form>
      </div>
    </div>
  );
}           