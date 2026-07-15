export type UserRole = "admin" | "user";

export interface AuthUser {
  id: string;
  name: string;
  email: string;
  role: UserRole;
}

export interface AuthSession {
  user: AuthUser;
  accessToken: string;
}