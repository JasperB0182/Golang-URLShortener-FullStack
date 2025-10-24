export interface UserRole {
  ID: number;
  Role: string;
}


export interface URLItem {
  ID: number;
  UserID: number;
  User: any;
  ShortCode: string;
  FullURL: string;
  CreatedAt: string;
  Enabled: boolean;
  ExpiryDate: string;
  UsageCount: number;
}

export interface URLListResponse {
  Code: URLItem[];
}
