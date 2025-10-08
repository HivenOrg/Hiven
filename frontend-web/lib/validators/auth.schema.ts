import {z} from "zod"

export const registerUserSchema = z.object({
    email: z.email(),
    password: z.string().min(6),
    firstname: z.string().min(1),
    lastname: z.string().min(1),
    phone_number: z.string().regex(/^\+\d{10,15}$/, "Phone number must be in E.164 format"),
})

export type RegisterUserInput = z.infer<typeof registerUserSchema>

export const loginUserSchema = z.object({
    email: z.email(),
    password: z.string().min(6),
})

export type LoginUserInput = z.infer<typeof loginUserSchema>
