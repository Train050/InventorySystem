import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS} from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { LandingPageComponent } from './pages/landing-page/landing-page.component';
import { InventoryHomePageComponent, AddItemComponent } from './pages/inventory-home-page/inventory-home-page.component';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule} from '@angular/material/menu';
import { MatTableModule} from '@angular/material/table';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { JwtInterceptor } from './jwt.interceptor';
import { MatFormFieldModule} from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatPaginatorModule} from '@angular/material/paginator';
import { MatDialogModule} from '@angular/material/dialog';
import { MatSortModule } from '@angular/material/sort';
import { MatIconModule} from '@angular/material/icon';
@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    LandingPageComponent,
    InventoryHomePageComponent,
    RegistrationPageComponent,
    ProfileComponent,
    AddItemComponent
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
    MatInputModule,
    MatPaginatorModule,
    MatDialogModule,
    ReactiveFormsModule,
    MatSortModule,
    MatIconModule,
  ],
  providers: [HttpClientModule,{provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true}],
  bootstrap: [AppComponent],
  schemas: [
    NO_ERRORS_SCHEMA
  ]
})
export class AppModule {}
