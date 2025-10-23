import { zodResolver } from "@hookform/resolvers/zod";
import { Eye } from "lucide-react";
import type { Metadata } from "next";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { useAuth } from "@/components/contexts/AuthContext";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  type RegisterUserInput,
  registerUserSchema,
} from "@/lib/validators/auth.schema";

export const metadata: Metadata = {
  title: "Sign Up",
  description: "Create your account",
};

export default function SignupPage() {
  const auth = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterUserInput>({
    resolver: zodResolver(registerUserSchema),
  });

  function onSubmit(data: RegisterUserInput) {
    auth.register(data);
  }
  return (
    <div className="flex min-h-screen flex-col bg-background p-4 pt-10">
      {/* Header with back button */}
      {/* <div className="flex items-center justify-between mb-8 pt-4">
        <Link href="/login" className="text-primary">
          <ArrowLeft size={24} />
        </Link>
      </div> */}

      {/* Title */}
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-primary">Create account</h1>
      </div>

      {/* Form */}
      <form onSubmit={handleSubmit(onSubmit)} className="flex-1 space-y-6">
        {/* Name fields - side by side */}
        <div className="grid grid-cols-2 gap-4">
          <div>
            <Label
              htmlFor="firstName"
              className="mb-2 block text-sm font-medium text-muted-foreground"
            >
              First name
            </Label>
            <Input
              id="firstName"
              placeholder="Liam"
              className="bg-card border-none rounded-lg text-card-foreground"
              {...register("firstname")}
            />
            {errors.firstname && (
              <p className="mt-1 text-sm text-red-600">
                {errors.firstname.message}
              </p>
            )}
          </div>
          <div>
            <Label
              htmlFor="lastName"
              className="mb-2 block text-sm font-medium text-muted-foreground"
            >
              Last name
            </Label>
            <Input
              id="lastName"
              placeholder="Smith"
              className="bg-card border-none rounded-lg text-card-foreground"
              {...register("lastname")}
            />
            {errors.lastname && (
              <p className="mt-1 text-sm text-red-600">
                {errors.lastname.message}
              </p>
            )}
          </div>
        </div>

        {/* Phone field */}
        <div>
          <Label
            htmlFor="phone"
            className="mb-2 block text-sm font-medium text-muted-foreground"
          >
            Phone
          </Label>
          <Input
            id="phone"
            type="tel"
            placeholder="(555) 555-5555"
            className="bg-card border-none rounded-lg text-card-foreground"
            {...register("phone_number")}
          />
          {errors.phone_number && (
            <p className="mt-1 text-sm text-red-600">
              {errors.phone_number.message}
            </p>
          )}
        </div>

        {/* Email field */}
        <div>
          <Label
            htmlFor="email"
            className="mb-2 block text-sm font-medium text-muted-foreground"
          >
            Email
          </Label>
          <Input
            id="email"
            type="email"
            placeholder="liam.smith@example.com"
            className="bg-card border-none rounded-lg text-card-foreground"
            {...register("email")}
          />
          {errors.email && (
            <p className="mt-1 text-sm text-red-600">
              {errors.email.message}
            </p>
          )}
        </div>

        {/* Password field */}
        <div>
          <Label
            htmlFor="password"
            className="mb-2 block text-sm font-medium text-muted-foreground"
          >
            Password
          </Label>
          <div className="relative">
            <Input
              id="password"
              type="password"
              placeholder="•••••"
              className="bg-card border-none rounded-lg text-card-foreground pr-12"
              {...register("password")}
            />
            <button
              type="button"
              className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
            >
              <Eye size={20} />
            </button>
          </div>
          {errors.password && (
            <p className="mt-1 text-sm text-red-600">
              {errors.password.message}
            </p>
          )}
        </div>

        {/* Spacer to push button to bottom */}
        <div className="flex-1" />

        {/* Register button */}
        <div className="pb-8">
          <Button
            type="submit"
            className="w-full bg-primary text-primary-foreground hover:bg-primary/90 rounded-full py-3 font-semibold"
          >
            Register
          </Button>
          <Link
            href="/login"
            className="text-sm text-muted-foreground hover:text-foreground text-center w-full block mt-2"
          >
            Already have an account? Log in
          </Link>
        </div>
      </form>
    </div>
  );
}
