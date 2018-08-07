

import android.app.Activity;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.support.v4.view.MotionEventCompat;
import android.support.v4.view.ViewCompat;
import android.view.LayoutInflater;
import android.view.View;
import android.view.View.OnLongClickListener;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.AdapterView.OnItemClickListener;
import android.widget.BaseAdapter;
import android.widget.Button;
import android.widget.LinearLayout;
import android.widget.ListView;
import android.widget.RadioButton;
import android.widget.RadioGroup;
import android.widget.RadioGroup.OnCheckedChangeListener;
import android.widget.SeekBar;
import android.widget.SeekBar.OnSeekBarChangeListener;
import android.widget.Spinner;
import android.widget.TextView;
import android.widget.Toast;
import java.util.ArrayList;
import java.util.GregorianCalendar;

public class PlanActivity extends Activity 
{



    public void nada()
    {
        switch (answer[0]) 
        {
            case '\u0001':
                switch (answer[1]) 
                {
                    case 'R':
                        switch (answer[2]) 
                        {
                            case 'F':
                                PlanActivity.this.pendingElements = (byte) answer[3];
                                if (PlanActivity.this.pendingElements > (byte) 0) 
                                {
                                    Toast.makeText(PlanActivity.this.getApplicationContext(), PlanActivity.this.getResources().getString(C0182R.string.eventAdd1) + " " + String.valueOf(PlanActivity.this.pendingElements) + " " + PlanActivity.this.getResources().getString(C0182R.string.eventAdd2), 0).show();
                                } 
                                else 
                                {
                                    PlanActivity.this.pendingElements = (byte) 0;
                                    Toast.makeText(PlanActivity.this.getApplicationContext(), PlanActivity.this.getResources().getString(C0182R.string.eventMemoryFull), 0).show();
                                }
                                PlanActivity.this.listEvents();
                            break;

                            case 'G':
                                String dayWeek = "";
                                switch (answer[3]) 
                                {
                                    case '\u0001':
                                        dayWeek = "Domingo";
                                    break;
                                    case '\u0002':
                                        dayWeek = "Lunes";
                                    break;
                                    case '\u0003':
                                        dayWeek = "Martes";
                                    break;
                                    case '\u0004':
                                        dayWeek = "Miercoles";
                                    break;
                                    case '\u0005':
                                        dayWeek = "Jueves";
                                    break;
                                    case '\u0006':
                                        dayWeek = "Viernes";
                                    break;
                                    case '\u0007':
                                        dayWeek = "Sabado";
                                    break;
                                }

                                int hora = PlanActivity.this.bcdToDecimal(answer[4] - 1);
                                int minuto = PlanActivity.this.bcdToDecimal(answer[5] - 1);
                                int segundo = PlanActivity.this.bcdToDecimal(answer[6] - 1);
                                PlanActivity.this.textViewTime.setText(dayWeek + ", " + String.format("%02d", new Object[]{Integer.valueOf(hora)}) + ":" + String.format("%02d", new Object[]{Integer.valueOf(minuto)}) + ":" + String.format("%02d", new Object[]{Integer.valueOf(segundo)}));
                            break;

                            case 'L':
                                Model model = new Model();
                                String[] daysOfWeek = PlanActivity.this.getResources().getStringArray(C0182R.array.daysOfWeeks);
                                if (answer[3] != '\u0001') 
                                {
                                    model.setFrecuency(daysOfWeek[answer[3] - 1]);
                                } 
                                else 
                                {
                                    model.setFrecuency(PlanActivity.this.getResources().getString(C0182R.string.daily));
                                }
                                model.setGroup(PlanActivity.this.getResources().getStringArray(C0182R.array.groups)[((answer[8] - 3) / 2) - 1]);
                                model.setTime(String.format("%d:%02d", new Object[]{Integer.valueOf(answer[4] - 1), Integer.valueOf(answer[5] - 1)}));
                                if (answer[9] == '\u000b') 
                                {
                                    int dimado = 100 - ((int) (3.2258065f * ((float) (((answer[10] - 3) - 2) / 2))));
                                    if (((answer[11] - 3) - 2) / 2 >= 0) 
                                    {
                                        model.setType(PlanActivity.this.getResources().getString(C0182R.string.dimmingColor));
                                        model.setDescription(String.format("%d%%-%d", new Object[]{Integer.valueOf(dimado), Integer.valueOf(color)}));
                                    } 
                                    else 
                                    {
                                        model.setType(PlanActivity.this.getResources().getString(C0182R.string.dimming));
                                        model.setDescription(String.format("%d%%", new Object[]{Integer.valueOf(dimado)}));
                                    }
                                } 
                                else if (answer[9] == '\t') 
                                {
                                    int indiceLight;
                                    model.setType(PlanActivity.this.getResources().getString(C0182R.string.sensor));
                                    String[] onTimes = PlanActivity.this.getResources().getStringArray(C0182R.array.onTimes);
                                    String[] offTimes = PlanActivity.this.getResources().getStringArray(C0182R.array.offTimes);
                                    String[] motionSensibilities = PlanActivity.this.getResources().getStringArray(C0182R.array.motionSensibilities);
                                    String[] lightSensibilities = PlanActivity.this.getResources().getStringArray(C0182R.array.lightSensibilities);
                                    if (answer[15] == '\u0013') 
                                    {
                                        indiceLight = 0;
                                    } 
                                    else 
                                    {
                                        indiceLight = (answer[15] - 3) / 2;
                                    }
                                    model.setDescription(String.format("%d%%-%s-%d%%-%s-%s-%s", new Object[]{Integer.valueOf(100 - ((int) (3.2258065f * ((float) (((answer[10] - 3) - 2) / 2))))), onTimes[((answer[11] - 3) - 2) / 2], Integer.valueOf(100 - ((int) (3.2258065f * ((float) (((answer[12] - 3) - 2) / 2))))), offTimes[((answer[13] - 3) - 2) / 2], motionSensibilities[7 - (((answer[14] - 3) - 2) / 2)], lightSensibilities[indiceLight]}));
                                }
                                PlanActivity.this.eventList.add(model);
                                PlanActivity.this.adapter.notifyDataSetChanged();
                                if (PlanActivity.access$3606(PlanActivity.this) <= (byte) 0) 
                                {
                                    PlanActivity.this.buttonNewEvent.setEnabled(true);
                                    break;
                                }
                            break;

                            case 'R':
                                Toast.makeText(PlanActivity.this.getApplicationContext(), PlanActivity.this.getResources().getString(C0182R.string.eventRemoved1) + " " + String.valueOf(answer[3]) + " " + PlanActivity.this.getResources().getString(C0182R.string.eventRemoved2), 0).show();
                                PlanActivity.this.pendingElements = (byte) answer[3];
                                if (PlanActivity.this.pendingElements == (byte) 0)
                                {
                                    PlanActivity.this.buttonNewEvent.setEnabled(true);
                                }
                                PlanActivity.this.listEvents();
                            break;

                            case 'S':
                            break;

                            default:
                            break;
                        }

                    default:

                    break;
                }

            case 'D':
                int indexStart = 1;
                int length = str.length();
                int indexEnd = str.indexOf(" HWV");
                if (indexEnd == -1 || indexEnd >= length) 
                {
                    PlanActivity.this.textViewDevice.setText("---");
                    PlanActivity.this.textViewHWV.setText("---");
                } 
                else 
                {
                    PlanActivity.this.textViewDevice.setText(str.substring(1, indexEnd));
                    indexStart = indexEnd + " HWV".length();
                }
                indexEnd = str.indexOf(" SWV");
                if (indexEnd == -1 || indexEnd >= length) 
                {
                    PlanActivity.this.textViewHWV.setText("---");
                    PlanActivity.this.textViewSWV.setText("---");
                } 
                else 
                {
                    PlanActivity.this.textViewHWV.setText(str.substring(indexStart, indexEnd));
                    indexStart = indexEnd + " SWV".length();
                    indexEnd = str.indexOf(0);
                    if (indexStart < indexEnd) 
                    {
                        PlanActivity.this.textViewSWV.setText(str.substring(indexStart, indexEnd));
                    }
                }
                if (!PlanActivity.this.textViewDevice.getText().equals("MPFAEXBT")) 
                {
                    Toast.makeText(PlanActivity.this, PlanActivity.this.getResources().getString(C0182R.string.incompatibleDevice), 0).show();
                    break;
                }

                PlanActivity.this.radioButtonDimming.setEnabled(true);
                PlanActivity.this.radioButtonSensor.setEnabled(true);
                PlanActivity.this.radioButtonDaily.setEnabled(true);
                PlanActivity.this.radioButtonWeekly.setEnabled(true);
                PlanActivity.this.spinnerDayOfWeek.setEnabled(true);
                PlanActivity.this.spinnerGroup.setEnabled(true);
                PlanActivity.this.spinnerHour.setEnabled(true);
                PlanActivity.this.spinnerMinute.setEnabled(true);
                PlanActivity.this.seekBarDimmingLevel.setEnabled(true);
                if (PlanActivity.this.multiColor) 
                {
                    PlanActivity.this.seekBarColorLevel.setEnabled(true);
                }
                PlanActivity.this.seekBarSensorLevelOn.setEnabled(true);
                PlanActivity.this.seekBarSensorMotion.setEnabled(true);
                if (PlanActivity.this.seekBarSensorMotion.getProgress() != 0) 
                {
                    PlanActivity.this.seekBarSensorTimeDimOn.setEnabled(true);
                    PlanActivity.this.seekBarSensorDimOff.setEnabled(true);
                    PlanActivity.this.seekBarSensorTimeDimOff.setEnabled(true);
                    PlanActivity.this.seekBarSensorLight.setEnabled(true);
                }
                PlanActivity.this.buttonNewEvent.setEnabled(true);
                PlanActivity.this.listEvents();
                break;
            break;
        }
    }


}