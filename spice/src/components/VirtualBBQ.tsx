import { Button, ButtonGroup, Grid, Paper, Slider, Typography } from '@mui/material';
import { Session } from 'bbq/intake/v1/bbq_pb';
import { useEffect, useState } from 'react';

interface SensorProbe {
  number: number;
  maxTemp: number;
  minTemp: number;
}

interface VirtualBBQProps {
  probes: SensorProbe[];
  session: Session | undefined;
  onTemperatureChange: (temperatures: number[]) => void;
  cooking: boolean;
  startCooking: () => void;
  stopCooking: () => void;
}

const VirtualBBQ = ({
  probes,
  session,
  onTemperatureChange,
  cooking,
  startCooking,
  stopCooking,
}: VirtualBBQProps) => {
  const [temperatures, setTemperature] = useState<number[]>([125, 125, 125, 125]);
  const valuetext = (value: number) => `${value}°F`;

  useEffect(() => {
    onTemperatureChange(temperatures);
  }, [temperatures, onTemperatureChange]);

  return (
    <Paper elevation={2} sx={{ p: 2, m: 2 }}>
      <h2>Virtual BBQ</h2>
      <Grid container>
        {probes.map((probe) => (
          <>
            <Grid item sm={2}>
              <Typography variant="h6">
                Probe {probe.number} {temperatures[probe.number]}°F
              </Typography>
            </Grid>
            <Grid item sm={10}>
              <Slider
                aria-label="Temperature"
                defaultValue={125}
                getAriaValueText={valuetext}
                valueLabelDisplay="auto"
                value={temperatures[probe.number]}
                shiftStep={5}
                step={5}
                min={probe.minTemp}
                max={probe.maxTemp}
                disabled={session === undefined}
                onChange={(e, value) => {
                  setTemperature((prev) => {
                    const next = [...prev];
                    next[probe.number] = value as number;
                    return next;
                  });
                }}
              />
            </Grid>
          </>
        ))}
        <Grid item sm={12} sx={{ textAlign: 'right' }}>
          <ButtonGroup>
            <Button
              variant="contained"
              onClick={startCooking}
              disabled={session === undefined || cooking}
            >
              Start Cooking
            </Button>
            <Button
              variant="contained"
              onClick={stopCooking}
              disabled={session === undefined || !cooking}
            >
              Stop Cooking
            </Button>
          </ButtonGroup>
        </Grid>
      </Grid>
    </Paper>
  );
};

export { VirtualBBQ, SensorProbe };
