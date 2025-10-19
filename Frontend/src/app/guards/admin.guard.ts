import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { map } from 'rxjs/operators';

export const adminGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);
  const router = inject(Router);

  return authService.checkAdminStatus().pipe(
    map(isAdmin => {
      if (!isAdmin) {
        router.navigate(['/login']);
        return false;
      }
      return true;
    })
  );
};
