import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RedirectService} from "../services/redirect.service";

@Component({
  selector: 'app-redirect',
  standalone: true,
  imports: [],
  templateUrl: './redirect.component.html',
  styleUrl: './redirect.component.scss'
})
export class RedirectComponent implements OnInit {
  protected id: string | null = "";

  constructor(private route: ActivatedRoute) {}

  protected redirectService = inject(RedirectService)

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
    if (this.id) {
      this.redirectService.redirect(this.id).subscribe({
        next: (res) => {
          window.location.href = res.URL
      },
        error: (err) => {
          window.location.href = "google.com"
        }
      })
    }

  }
}
