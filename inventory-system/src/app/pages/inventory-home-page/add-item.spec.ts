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
import { AddItemComponent } from './add-item';

describe('AddItemComponent', () => {
  let component: AddItemComponent;
  let fixture: ComponentFixture<AddItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MatFormFieldModule, MatGridListModule, MatPaginatorModule, FormsModule, MatInputModule, MatSortModule, BrowserAnimationsModule, MatDialogModule],
      declarations: [ AddItemComponent ],
      schemas: [NO_ERRORS_SCHEMA]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  beforeEach(() =>{
    fixture = TestBed.createComponent(AddItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges;
  })

  it('button click event one way', async(() => {
    spyOn(component, 'submitItem');

    let button = fixture.debugElement.nativeElement.querySelector('button');
    button.click(); // you can use     btn.triggerEventHandler('click', null);

    fixture.detectChanges();

    fixture.whenStable().then(() => {
      expect(component.submitItem).toHaveBeenCalled();
    });
  }));
});


