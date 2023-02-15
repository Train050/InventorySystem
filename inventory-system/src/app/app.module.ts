import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {HttpClientModule, HTTP_INTERCEPTORS} from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { LandingPageComponent } from './pages/landing-page/landing-page.component';
import { InventoryHomePageComponent } from './pages/inventory-home-page/inventory-home-page.component';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule} from '@angular/material/menu';
import { MatTableModule} from '@angular/material/table';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { JwtInterceptor } from './jwt.interceptor';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatPaginatorModule} from '@angular/material/paginator';
@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    LandingPageComponent,
    InventoryHomePageComponent,
    RegistrationPageComponent,
    ProfileComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    MatGridListModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatMenuModule,
    HttpClientModule,
    FormsModule,
    MatTableModule,
    MatFormFieldModule,
    MatPaginatorModule
  ],
  providers: [{provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true}],
  bootstrap: [AppComponent],
})
export class AppModule {}
