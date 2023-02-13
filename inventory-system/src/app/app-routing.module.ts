import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth.guard';
import { ProfileComponent } from './pages/profile/profile.component';
import { LandingPageComponent } from './pages/landing-page/landing-page.component';
import { InventoryHomePageComponent } from './pages/inventory-home-page/inventory-home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';

const routes: Routes = [
  { path: 'landing-page', component: LandingPageComponent },
  {
    path: 'inventory-home-page',
    component: InventoryHomePageComponent,
  },
  { path: 'login-page', component: LoginPageComponent },
  {path: '', redirectTo: '/profile', pathMatch: 'full'},
  {
    path: 'profile',
    component: ProfileComponent,
    canActivate: [AuthGuard]
  },
  { path: 'registration-page', component: RegistrationPageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
