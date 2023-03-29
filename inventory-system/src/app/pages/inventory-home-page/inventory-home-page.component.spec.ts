import { ComponentFixture, TestBed, async, fakeAsync, tick } from '@angular/core/testing';
import { InventoryHomePageComponent } from './inventory-home-page.component';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatInputModule } from '@angular/material/input';
import { MatPaginatorModule, MatPaginator } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { MatDialogModule } from '@angular/material/dialog';
import { fn } from 'cypress/types/jquery';

describe('InventoryHomePageComponent', () => {
  let component: InventoryHomePageComponent;
  let fixture: ComponentFixture<InventoryHomePageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MatFormFieldModule, MatGridListModule, MatPaginatorModule, FormsModule, MatInputModule, MatSortModule, BrowserAnimationsModule, MatDialogModule],
      declarations: [ InventoryHomePageComponent ],
      schemas: [NO_ERRORS_SCHEMA]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InventoryHomePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    expect(component.dataSource).toBeTruthy();
    expect(component.paginator).toBeInstanceOf(MatPaginator);
    expect(component.applyFilter).toBeTruthy();
    expect(component).toBeDefined();
    expect(component.openAddItem).toBeTruthy();
  });

});


