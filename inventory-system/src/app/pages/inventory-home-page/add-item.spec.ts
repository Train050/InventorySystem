import { ComponentFixture, TestBed, async, fakeAsync, tick } from '@angular/core/testing';
import { AddItemComponent } from './add-item';
import { MatDialogModule, MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { By } from '@angular/platform-browser';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('AddItemComponent', () => {
  let component: AddItemComponent;
  let fixture: ComponentFixture<AddItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MatDialogModule, HttpClientTestingModule, MatFormFieldModule, ReactiveFormsModule, MatInputModule, BrowserAnimationsModule],
      declarations: [ AddItemComponent ],
      providers: [
        { provide: MAT_DIALOG_DATA, useValue: {} },
        { provide: MatDialogRef, useValue: {} }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('button click event one way', async(() => {
    spyOn(component, 'submitItem');

    let button = fixture.debugElement.nativeElement.querySelector('button');
    button.click(); // you can use     btn.triggerEventHandler('click', null);

    fixture.detectChanges();

    fixture.whenStable().then(() => {
      expect(component.submitItem).toHaveBeenCalled();
    });
  }));

    it('should call submit() method on form submit', () => {
      /*Get button from html*/
      fixture.detectChanges();
      const compiled = fixture.debugElement.nativeElement;
      // Supply id of your form below formID
      const getForm = fixture.debugElement.query(By.css('FormGroup'));
      expect(getForm.triggerEventHandler('submit')).toBeTrue();
    });
});

