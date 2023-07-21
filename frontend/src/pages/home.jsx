import './home.css'
import { IonCard, IonCardContent, IonCardHeader, IonCardSubtitle, IonCardTitle } from '@ionic/react';

const Home = () => {
    return (
        <div>
        <div className='header'>
          <p>
            Tools To Protect Creative Work from Generative AI
          </p>
        </div>
        <IonCard class="custom-card">
      <IonCardHeader>
        <IonCardTitle>Card Title</IonCardTitle>
        <IonCardSubtitle>Card Subtitle</IonCardSubtitle>
      </IonCardHeader>

      <IonCardContent>Here's a small text description for the card content. Nothing more, nothing less.</IonCardContent>
    </IonCard>
      </div>
    );
  };
  
export default Home;