export interface UserRole {
  ID: number;
  Role: string;
}

export interface User {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  Email: string;
  Password: string;
  Name: string;
  RoleID: number;
  UserRole: UserRole;
}

export interface UsersResponse {
  Users: User[];
}
