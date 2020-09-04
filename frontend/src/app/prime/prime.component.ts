import { Component, OnInit } from '@angular/core';
import { PrimeService } from './prime.service';
import { Prime } from './prime';

@Component({
  selector: 'app-prime',
  providers: [PrimeService],
  templateUrl: 'prime.component.html',
  styleUrls: ['./prime.component.css']
})
export class PrimeComponent implements OnInit {
  prime: Prime;
  error: String;
  constructor(private primeService: PrimeService) { }
  ngOnInit(): void {
  }
  getPrime(no: string): void {
    if (this.validateNumber(no)) {
      this.primeService.getPrime(no).subscribe(prime => (this.prime = prime));
      this.error = undefined;
    } else {
      this.prime = undefined;
      this.error = "input must be number and in range [0, 2147483647]";
    }
  }

  validateNumber(no: String): Boolean {
    let num = parseInt(no, 10);
    return num !== NaN && num >= 0 && num < 2147483648;
  }
}
